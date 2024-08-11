package adapter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"ecommerce/internal/client/shopify"
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
	//GetShopifyAuthConfig() *oauth2.Config
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

	var externalProducts []*model.ExternalVariant
	for _, product := range products.Products {
		for _, variant := range product.Variants {
			var externalProduct *model.ExternalVariant
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				Result:  &externalProduct,
				TagName: "shopify",
			})
			if err != nil {
				return nil, err
			}

			err = decoder.Decode(variant)
			if err != nil {
				return nil, err
			}

			var externalProductId = strconv.FormatInt(product.ID, 10)

			externalProduct.IDExternal = strconv.FormatInt(variant.ID, 10)
			externalProduct.IDExternalProduct = &externalProductId

			if foundImage, ok := lo.Find(product.Images, func(image *shopify.Image) bool {
				if variant.ImageID == nil {
					return false
				}
				return image.ID == *variant.ImageID
			}); ok {
				externalProduct.ImageURL = foundImage.Src
			}

			if price, err := strconv.ParseFloat(variant.Price, 64); err == nil {
				externalProduct.Price = &price
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

			externalProduct.Name = product.Title
			externalProduct.Option = pgtype.JSON{
				Bytes:  jsonData,
				Status: pgtype.Present,
			}

			externalProducts = append(externalProducts, externalProduct)
		}
	}
	return externalProducts, nil
}
