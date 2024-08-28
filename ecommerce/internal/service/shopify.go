package service

import (
	"errors"
	"fmt"
	"regexp"

	"ecommerce/internal/adapter"
	"ecommerce/internal/client/shopify"
	"ecommerce/internal/constants"
	"ecommerce/internal/database"
	entity "ecommerce/internal/database/model"
	"ecommerce/internal/database/table"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"

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
	//CreateExternalVariants(shopifyProductIdList interface{}) error

	ConnectNewExternalShopShopify(shopDomain string, authorizeCode string) error

	SyncVariants(externalShopId int64) error
	GetStockByExternalProductExternalId(externalShopId int64, externalProductIdMappings []string) ([]*model.ExternalVariantStock, error)
}

type ShopifyService struct {
	Database *database.PostgresqlDatabase

	ProductRepository          repository.IProductRepository
	VariantRepository          repository.IVariantRepository
	ExternalShopRepository     repository.IExternalShopRepository
	ExternalShopAuthRepository repository.IExternalShopShopifyAuthRepository
	ExternalVariantRepository  repository.IExternalVariantRepository

	ShopifyAdapter adapter.IShopifyAdapter
}

func NewShopifyService(
	shopifyAdapter adapter.IShopifyAdapter,
	productRepository repository.IProductRepository,
	variantRepository repository.IVariantRepository,
	externalShopRepository repository.IExternalShopRepository,
	externalShopAuthRepository repository.IExternalShopShopifyAuthRepository,
	ExternalVariant repository.IExternalVariantRepository) IShopifyService {
	return &ShopifyService{
		ProductRepository:          productRepository,
		VariantRepository:          variantRepository,
		ExternalShopRepository:     externalShopRepository,
		ExternalShopAuthRepository: externalShopAuthRepository,
		ExternalVariantRepository:  ExternalVariant,
		ShopifyAdapter:             shopifyAdapter,
	}
}

func (s *ShopifyService) GetEcommerceId() int16 {
	return constants.SHOPIFY_ID
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
			entity.ExternalShop{
				Name:        shopName,
				FkShop:      1,
				FkEcommerce: constants.SHOPIFY_ID,
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
			entity.ExternalShopShopifyAuth{
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

func (s *ShopifyService) SyncVariants(externalShopId int64) error {
	externalShopAuth, err := s.ExternalShopAuthRepository.GetByExternalShopId(s.ExternalShopRepository.GetDatabase().Db, externalShopId)
	if err != nil {
		return err
	}

	externalVariants, err := s.ShopifyAdapter.GetProducts(&shopify.ShopifyClientParam{
		ShopName:    externalShopAuth.Name,
		AccessToken: *externalShopAuth.AccessToken,
	})
	if err != nil {
		return err
	}

	entityExternalVariants := lop.Map(externalVariants, func(variant *model.ExternalVariant, _ int) *entity.ExternalVariant {
		return &entity.ExternalVariant{
			FkExternalShop:           externalShopId,
			ExternalProductIDMapping: variant.ExternalProductIdMapping,
			ExternalIDMapping:        variant.ExternalIdMapping,
			Sku:                      variant.Sku,
			Name:                     variant.Name,
			Option:                   variant.Option,
			Status:                   variant.Status,
			Price:                    variant.Price,
			ImageURL:                 &variant.ImageUrl,
		}
	})

	_, err = s.ExternalVariantRepository.CreateMany(
		s.ExternalVariantRepository.GetDatabase().Db,
		postgres.ColumnList{
			table.ExternalVariant.FkExternalShop,
			table.ExternalVariant.ExternalProductIDMapping,
			table.ExternalVariant.ExternalIDMapping,
			table.ExternalVariant.Sku,
			table.ExternalVariant.Name,
			table.ExternalVariant.Option,
			table.ExternalVariant.Status,
			table.ExternalVariant.Price,
			table.ExternalVariant.ImageURL,
		},
		entityExternalVariants,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *ShopifyService) GetStockByExternalProductExternalId(externalShopId int64, externalProductIdMappings []string) ([]*model.ExternalVariantStock, error) {
	externalShopAuth, err := s.ExternalShopAuthRepository.GetByExternalShopId(s.ExternalShopRepository.GetDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	stocks, err := s.ShopifyAdapter.GetExternalVariantStockByproductIds(&shopify.ShopifyClientParam{
		ShopName:    externalShopAuth.Name,
		AccessToken: *externalShopAuth.AccessToken,
	}, externalProductIdMappings)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
