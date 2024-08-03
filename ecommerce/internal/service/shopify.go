package service

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"ecommerce/internal/adapter"
	"ecommerce/internal/client/shopify"
	"ecommerce/internal/constants"
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	lop "github.com/samber/lo/parallel"
)

const (
	ShopifyShopRegex = `^([a-zA-Z0-9][a-zA-Z0-9\-]*)\.myshopify\.com$`
)

type IShopifyService interface {
	GetEcommerceId() int16
	GetAuthorizePath(shopDomain string) string
	CreateProductVariants(shopifyProductIdList interface{}) error

	ConnectNewExternalShopShopify(shopDomain string, authorizeCode string) error
	SyncProducts(externalShopId int64) error
	GetExternalProductsByExternalShopId(externalShopId int64, limit int32, offset int32) (interface{}, error)
}

type ShopifyService struct {
	Database *database.PostgresqlDatabase

	ProductRepository                repository.IProductRepository
	VariantRepository                repository.IVariantRepository
	ExternalShopRepository           repository.IExternalShopRepository
	ExternalShopAuthRepository       repository.IExternalShopShopifyAuthRepository
	ExternalProductShopifyRepository repository.IExternalProductShopifyRepository

	ShopifyAdapter adapter.IShopifyAdapter
}

func NewShopifyService(
	shopifyAdapter adapter.IShopifyAdapter,
	productRepository repository.IProductRepository,
	variantRepository repository.IVariantRepository,
	externalShopRepository repository.IExternalShopRepository,
	externalShopAuthRepository repository.IExternalShopShopifyAuthRepository,
	externalProductShopify repository.IExternalProductShopifyRepository) IShopifyService {
	return &ShopifyService{
		ProductRepository:                productRepository,
		VariantRepository:                variantRepository,
		ExternalShopRepository:           externalShopRepository,
		ExternalShopAuthRepository:       externalShopAuthRepository,
		ExternalProductShopifyRepository: externalProductShopify,
		ShopifyAdapter:                   shopifyAdapter,
	}
}

func (s *ShopifyService) GetEcommerceId() int16 {
	return constants.SHOPIFY
}

func (s *ShopifyService) getShopOriginFromShopDomain(shopDomain string) string {
	regex := regexp.MustCompile(ShopifyShopRegex)
	matches := regex.FindStringSubmatch(shopDomain)
	return matches[1]
}

func (s *ShopifyService) GetAuthorizePath(shopDomain string) string {
	shopName := s.getShopOriginFromShopDomain(shopDomain)
	return s.ShopifyAdapter.GetAuthorizePath(&shopify.ShopifyClientParam{
		ShopName: shopName,
	})
}

func (s *ShopifyService) getAccessToken(shopDomain string, code string) (string, error) {
	shopName := s.getShopOriginFromShopDomain(shopDomain)
	accessToken, err := s.ShopifyAdapter.GetAccessToken(&shopify.ShopifyClientParam{
		ShopName: shopName,
	}, code)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *ShopifyService) ConnectNewExternalShopShopify(shopDomain string, authorizeCode string) error {
	accessToken, err := s.getAccessToken(shopDomain, authorizeCode)
	if err != nil {
		return fmt.Errorf("failed to get access token: %v", err)
	}

	shopName := s.getShopOriginFromShopDomain(shopDomain)
	return s.createExternalShopShopify(shopName, accessToken)
}

func (s *ShopifyService) createExternalShopShopify(shopName string, accessToken string) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		newExternalShop, err := s.ExternalShopRepository.CreateOne(db,
			postgres.ColumnList{table.ExternalShop.Name, table.ExternalShop.FkShop, table.ExternalShop.FkEcommerce},
			model.ExternalShop{
				Name:        shopName,
				FkShop:      1,
				FkEcommerce: constants.SHOPIFY,
			})
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("user already connected to this Shopify store")
			} else {
				return nil, fmt.Errorf("failed to create new Shopify external shop: %v", err)
			}
		}

		if err != nil {
			return nil, fmt.Errorf("failed to create new Shopify external shop: %v", err)
		}

		if _, err := s.ExternalShopAuthRepository.CreateOne(db,
			postgres.ColumnList{
				table.ExternalShopShopifyAuth.FkExternalShop,
				table.ExternalShopShopifyAuth.Name,
				table.ExternalShopShopifyAuth.AccessToken,
			},
			model.ExternalShopShopifyAuth{
				FkExternalShop: newExternalShop.IDExternalShop,
				Name:           shopName,
				AccessToken:    &accessToken,
			}); err != nil {
			return nil, fmt.Errorf("failed to create new Shopify auth: %v", err)
		}
		return nil, nil
	}

	_, err := s.ExternalShopRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (s *ShopifyService) SyncProducts(externalShopId int64) error {
	externalShopAuth, err := s.ExternalShopAuthRepository.GetByExternalShopId(s.ExternalShopRepository.GetDefaultDatabase().Db, externalShopId)
	if err != nil {
		return err
	}

	externalProducts, err := s.ShopifyAdapter.GetProducts(&shopify.ShopifyClientParam{
		ShopName:    externalShopAuth.Name,
		AccessToken: *externalShopAuth.AccessToken,
	})
	if err != nil {
		return err
	}

	externalProducts = lop.Map(externalProducts, func(product *model.ExternalProductShopify, _ int) *model.ExternalProductShopify {
		product.FkExternalShop = externalShopId
		return product
	})

	_, err = s.ExternalProductShopifyRepository.CreateMany(
		s.ExternalProductShopifyRepository.GetDefaultDatabase().Db,
		postgres.ColumnList{
			table.ExternalProductShopify.FkExternalShop,
			table.ExternalProductShopify.ShopifyProductID,
			table.ExternalProductShopify.ShopifyVariantID,
			table.ExternalProductShopify.Name,
			table.ExternalProductShopify.Sku,
			table.ExternalProductShopify.Stock,
			table.ExternalProductShopify.Option,
			table.ExternalProductShopify.Price,
			table.ExternalProductShopify.ImageURL,
		},
		externalProducts,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ShopifyService) GetExternalProductsByExternalShopId(externalShopId int64, limit int32, offset int32) (interface{}, error) {
	externalProducts, err := s.ExternalProductShopifyRepository.GetByExternalShopId(s.ExternalProductShopifyRepository.GetDefaultDatabase().Db, externalShopId, limit, offset)
	if err != nil {
		return nil, err
	}
	return externalProducts, nil
}

func (s *ShopifyService) CreateProductVariants(shopifyProductIdList interface{}) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		//var updateProductVariantList [][]int64

		//shouldUpdateProductVariants := postgres.CTE("should_update_product_variants")

		for _, shopifyProductId := range shopifyProductIdList.([]string) {
			_shopifyProductId, err := strconv.ParseInt(shopifyProductId, 10, 64)
			if err != nil {
				return nil, err
			}

			externalProducts, err := s.ExternalProductShopifyRepository.GetByShopifyProductId(s.ExternalProductShopifyRepository.GetDefaultDatabase().Db, int64(_shopifyProductId))
			if err != nil {
				return nil, err
			}

			newProduct, err := s.ProductRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.Product.Name,
					table.Product.FkShop,
				},
				model.Product{
					Name:   externalProducts[0].Name,
					FkShop: 1,
				},
			)
			if err != nil {
				return nil, err
			}

			for _, externalProduct := range externalProducts {
				newVariant, err := s.VariantRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.Variant.FkProduct,
						table.Variant.Name,
						table.Variant.Sku,
						table.Variant.Stock,
						table.Variant.Price,
					},
					model.Variant{
						FkProduct: newProduct.IDProduct,
						Name:      externalProduct.Name,
						Sku:       externalProduct.Sku,
						Stock:     externalProduct.Stock,
						Price:     externalProduct.Price,
					},
				)
				if err != nil {
					return nil, err
				}

				//updateProductVariantList = append(updateProductVariantList, []int64{newProduct.IDProduct, newVariant.IDVariant, externalProduct.ShopifyVariantID})

				if err := s.ExternalProductShopifyRepository.UpdateProductVariant(db, newProduct.IDProduct, newVariant.IDVariant, externalProduct.ShopifyVariantID); err != nil {
					return nil, err
				}
			}
		}

		return nil, nil
	}

	if _, err := s.ProductRepository.ExecWithinTransaction(execWithinTransaction); err != nil {
		return err
	}

	return nil
}
