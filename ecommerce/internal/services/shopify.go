package services

import (
	"errors"
	"fmt"
	"regexp"

	"ecommerce/internal/adapters"
	"ecommerce/internal/clients/shopify"
	"ecommerce/internal/configs"
	"ecommerce/internal/constants"
	"ecommerce/internal/database"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
)

const (
	ShopifyShopRegex = `^([a-zA-Z0-9][a-zA-Z0-9\-]*)\.myshopify\.com$`
)

type IShopifyService interface {
	GetEcommerceId() int32
	GetAuthorizePath(shopDomain string) string

	ConnectNewShopifyExternalShop(shopDomain string, authorizeCode string) error
	SyncProducts(externalShopId int32) (interface{}, error)
}

type ShopifyService struct {
	Config *configs.Config
	//Repository repositories.IShopifyRepository
	Database *database.PostgresqlDatabase

	ExternalShopRepository     repositories.IExternalShopRepository
	ExternalShopAuthRepository repositories.IShopifyExternalShopAuthRepository

	ShopifyAdapter adapters.IShopifyAdapter
}

func NewShopifyService(shopifyAdapter adapters.IShopifyAdapter, config *configs.Config, externalShopRepository repositories.IExternalShopRepository, externalShopAuthRepository repositories.IShopifyExternalShopAuthRepository) IShopifyService {
	return &ShopifyService{
		Config:                     config,
		ExternalShopRepository:     externalShopRepository,
		ExternalShopAuthRepository: externalShopAuthRepository,
		ShopifyAdapter:             shopifyAdapter,
	}
}

func (s *ShopifyService) GetEcommerceId() int32 {
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

func (s *ShopifyService) ConnectNewShopifyExternalShop(shopDomain string, authorizeCode string) error {
	accessToken, err := s.getAccessToken(shopDomain, authorizeCode)
	if err != nil {
		return fmt.Errorf("failed to get access token: %v", err)
	}

	shopName := s.getShopOriginFromShopDomain(shopDomain)
	return s.createShopifyExternalShop(shopName, accessToken)
}

func (s *ShopifyService) createShopifyExternalShop(shopName string, accessToken string) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		newExternalShop, err := s.ExternalShopRepository.Create(db, models.ExternalShop{
			Name:        shopName,
			FkEcommerce: constants.SHOPIFY,
		}, postgres.ColumnList{s.ExternalShopRepository.GetTable().Name, s.ExternalShopRepository.GetTable().FkEcommerce})
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

		if _, err := s.ExternalShopAuthRepository.Create(db, models.ShopifyExternalShopAuth{
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

func (s *ShopifyService) SyncProducts(externalShopId int32) (interface{}, error) {
	externalShopAuth, err := s.ExternalShopAuthRepository.GetByExternalShopId(s.ExternalShopRepository.GetDefaultDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	products, err := s.ShopifyAdapter.GetProducts(&shopify.ShopifyClientParam{
		ShopName:    externalShopAuth.Name,
		AccessToken: *externalShopAuth.AccessToken,
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", products)
	return products, nil
}
