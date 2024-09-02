package adapter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"ecommerce/internal/client/shopify"
	clientModel "ecommerce/internal/client/shopify/model"
	"ecommerce/internal/configs"
	"ecommerce/internal/constants"
	"ecommerce/internal/model"

	"github.com/go-viper/mapstructure/v2"
	"github.com/jackc/pgtype"
	"github.com/samber/lo"
	"golang.org/x/oauth2"
)

const (
	ShopifyBaseURL         = "https://%s.myshopify.com"
	ShopifyTokenKey        = "X-Shopify-Access-Token"
	ShopifyAuthorizePath   = "/admin/oauth/authorize"
	ShopifyRedirectPath    = "/api/shopify/redirect"
	ShopifyAccessTokenPath = "/admin/oauth/access_token"
)

type IShopifyAdapter interface {
	createOauth2Config() *oauth2.Config
	getShopifyClient(param *shopify.ShopifyClientParam) shopify.IShopifyClient

	GetAuthorizePath(param *shopify.ShopifyClientParam) string
	GetAccessToken(param *shopify.ShopifyClientParam, code string) (string, error)
	GetProducts(param *shopify.ShopifyClientParam) ([]*model.ExternalVariant, error)
	GetExternalVariantStockByproductIds(param *shopify.ShopifyClientParam, productIds []string) ([]*model.ExternalVariantStock, error)

	CreateOrder(param *shopify.ShopifyClientParam, externalOrderItems []*model.ExternalOrderItem) (string, error)
}

type ShopifyAdapterConfig struct {
	ClientID     string
	ClientSecret string
	HostBaseUrl  string
}

type ShopifyAdapter struct {
	Config *ShopifyAdapterConfig
}

func NewShopifyAdapter(config *configs.Config) IShopifyAdapter {
	return &ShopifyAdapter{
		Config: &ShopifyAdapterConfig{
			ClientID:     config.ShopifyApp.ClientID,
			ClientSecret: config.ShopifyApp.ClientSecret,
			HostBaseUrl:  fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		},
	}
}

func (a *ShopifyAdapter) getShopifyClient(param *shopify.ShopifyClientParam) shopify.IShopifyClient {
	return shopify.NewShopifyClient(param)
}

func (a *ShopifyAdapter) createOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     a.Config.ClientID,
		ClientSecret: a.Config.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", a.Config.HostBaseUrl, constants.ShopifyRedirectPath),
		Scopes:       []string{"read_products", "read_orders", "write_orders"},
	}
}

func (a *ShopifyAdapter) GetAuthorizePath(param *shopify.ShopifyClientParam) string {
	oauth2Config := a.createOauth2Config()
	return a.getShopifyClient(param).GetAuthorizePath(oauth2Config)
}

func (a *ShopifyAdapter) GetAccessToken(param *shopify.ShopifyClientParam, code string) (string, error) {
	oauth2Config := a.createOauth2Config()
	token, err := a.getShopifyClient(param).GetAccessToken(oauth2Config, code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (a *ShopifyAdapter) GetProducts(param *shopify.ShopifyClientParam) ([]*model.ExternalVariant, error) {
	products, err := a.getShopifyClient(param).GetProducts()
	if err != nil {
		return nil, err
	}

	var externalVariants []*model.ExternalVariant
	for _, product := range products.Products {
		for _, variant := range product.Variants {
			var externalVariant *model.ExternalVariant
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				Result:  &externalVariant,
				TagName: "shopify",
			})
			if err != nil {
				return nil, err
			}

			err = decoder.Decode(variant)
			if err != nil {
				return nil, err
			}

			var externalProductIdMapping = strconv.FormatInt(product.ID, 10)

			externalVariant.ExternalIdMapping = strconv.FormatInt(variant.ID, 10)
			externalVariant.ExternalProductIdMapping = &externalProductIdMapping

			if foundImage, ok := lo.Find(product.Images, func(image *shopify.Image) bool {
				if variant.ImageID == nil {
					return false
				}
				return image.ID == *variant.ImageID
			}); ok {
				externalVariant.ImageUrl = *foundImage.Src
			} else {
				if product.Image != nil && product.Image.Src != nil {
					externalVariant.ImageUrl = *product.Image.Src
				}
			}

			if price, err := strconv.ParseFloat(variant.Price, 64); err == nil {
				externalVariant.Price = price
			}

			_option := make(map[string]string)
			r := reflect.ValueOf(variant)
			for index, option := range product.Options {
				indexName := fmt.Sprintf("Option%d", index+1)

				field := reflect.Indirect(r).FieldByName(indexName)
				_option[option.Name] = field.String()
			}

			jsonData, err := json.Marshal(_option)
			if err != nil {
				continue
			}

			externalVariant.Name = product.Title
			externalVariant.Sku = &variant.Sku
			if product.Status == constants.ACTIVE {
				externalVariant.Status = constants.ACTIVE
			} else {
				externalVariant.Status = constants.INACTIVE
			}
			externalVariant.Option = pgtype.JSON{
				Bytes:  jsonData,
				Status: pgtype.Present,
			}

			externalVariants = append(externalVariants, externalVariant)
		}
	}
	return externalVariants, nil
}

func (a *ShopifyAdapter) GetExternalVariantStockByproductIds(param *shopify.ShopifyClientParam, productIds []string) ([]*model.ExternalVariantStock, error) {
	externalProducts, err := a.getShopifyClient(param).GetProductsByProductIds(productIds)
	if err != nil {
		return nil, err
	}

	var externalVariantStocks []*model.ExternalVariantStock
	for _, product := range externalProducts.Products {
		for _, variant := range product.Variants {
			var externalProductIdMapping = strconv.FormatInt(product.ID, 10)

			externalVariantStocks = append(externalVariantStocks, &model.ExternalVariantStock{
				ExternalProductIdMapping: &externalProductIdMapping,
				ExternalIdMapping:        strconv.FormatInt(variant.ID, 10),
				Stock:                    variant.InventoryQuantity,
			})
		}
	}
	return externalVariantStocks, nil
}

func (a *ShopifyAdapter) CreateOrder(param *shopify.ShopifyClientParam, externalOrderItems []*model.ExternalOrderItem) (string, error) {
	createOrderRequest := &clientModel.CreateOrderRequest{
		Order: &clientModel.OrderRequest{
			LineItems: make([]*clientModel.LineItemRequest, 0),
		},
	}
	for _, externalOrderItem := range externalOrderItems {
		variantId, err := strconv.Atoi(externalOrderItem.ExternalIdMapping)
		if err != nil {
			return "", err
		}
		createOrderRequest.Order.LineItems = append(createOrderRequest.Order.LineItems, &clientModel.LineItemRequest{
			VariantID: variantId,
			Quantity:  externalOrderItem.Quantity,
		})

	}

	newExternalOrder, err := a.getShopifyClient(param).CreateOrder(createOrderRequest)
	if err != nil {
		return "", err
	}
	return string(newExternalOrder.Order.ID), nil
}
