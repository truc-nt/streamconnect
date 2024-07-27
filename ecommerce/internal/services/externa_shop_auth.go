package services

import (
	"ecommerce/internal/configs"
	"ecommerce/internal/constants"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"
	"errors"
)

type IExternalShopAuthService interface {
	getExternalShopAuth(ecommerceId int32) (interface{}, error)

	CreateShopifyAuth(externalShopId int32, accessToken string) (int32, error)
	GetShopifyAuthByExternalShopId(externalShopId int32) (*models.ShopifyExternalShopAuth, error)
}

type ExternalShopAuthService struct {
	Config *configs.Config

	ShopifyExternalShopAuthRepository repositories.IShopifyExternalShopAuthRepository
}

func NewExternalShopAuthService(repo repositories.IShopifyExternalShopAuthRepository, config *configs.Config) IExternalShopAuthService {
	return &ExternalShopAuthService{
		ShopifyExternalShopAuthRepository: repo,
		Config:                            config,
	}
}

func (s *ExternalShopAuthService) getExternalShopAuth(ecommerceId int32) (interface{}, error) {
	switch ecommerceId {
	case constants.SHOPIFY:
		return s.ShopifyExternalShopAuthRepository, nil
	default:
		break
	}
	return nil, errors.New("Ecommerce adapter not found")
}

func (s *ExternalShopAuthService) CreateShopifyAuth(externalShopId int32, accessToken string) (int32, error) {
	newData, err := s.ShopifyExternalShopAuthRepository.Create(s.ShopifyExternalShopAuthRepository.GetDefaultDatabase().Db, models.ShopifyExternalShopAuth{
		FkExternalShop: externalShopId,
		AccessToken:    &accessToken,
	})
	if err != nil {
		return 0, err
	}

	return newData.IDShopifyExternalShopAuth, nil
}

func (s *ExternalShopAuthService) GetShopifyAuthByExternalShopId(externalShopId int32) (*models.ShopifyExternalShopAuth, error) {
	auth, err := s.ShopifyExternalShopAuthRepository.GetByExternalShopId(s.ShopifyExternalShopAuthRepository.GetDefaultDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
