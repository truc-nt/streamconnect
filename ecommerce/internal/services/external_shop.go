package services

import (
	"ecommerce/internal/configs"
	"ecommerce/internal/constants"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"

	"github.com/go-jet/jet/v2/postgres"
)

type IExternalShopService interface {
	CreateShopifyExternalShop(name string) (int32, error)
	GetExternalShopsByShopId(shopId int32, limit int32, offset int32) ([]*models.ExternalShop, error)
	SyncProductsByExternalShopId(externalShopId int32) error
}

type ExternalShopService struct {
	Config     *configs.Config
	Repository repositories.IExternalShopRepository

	EcommerceService map[int32]IEcommerceService
}

func NewExternalShopService(repo repositories.IExternalShopRepository, config *configs.Config, ecommerceServices []IEcommerceService) IExternalShopService {
	ecommerceServicesMap := make(map[int32]IEcommerceService)
	for _, ecommerceService := range ecommerceServices {
		ecommerceServicesMap[ecommerceService.GetEcommerceId()] = ecommerceService
	}

	return &ExternalShopService{
		Repository:       repo,
		Config:           config,
		EcommerceService: ecommerceServicesMap,
	}
}

func (s *ExternalShopService) getEcommerceExternalShopService(ecommerceId int32) (IEcommerceService, error) {
	return s.EcommerceService[ecommerceId], nil
}

func (s *ExternalShopService) CreateShopifyExternalShop(name string) (int32, error) {
	newExternalShop := models.ExternalShop{
		Name:        name,
		FkEcommerce: constants.SHOPIFY,
	}

	newData, err := s.Repository.Create(s.Repository.GetDefaultDatabase().Db, newExternalShop, postgres.ColumnList{s.Repository.GetTable().Name, s.Repository.GetTable().FkEcommerce})
	if err != nil {
		return 0, err
	}

	return newData.IDExternalShop, nil
}

func (s *ExternalShopService) GetExternalShopsByShopId(shopId int32, limit int32, offset int32) ([]*models.ExternalShop, error) {
	shops, err := s.Repository.GetByShopId(s.Repository.GetDefaultDatabase().Db, shopId, limit, offset)
	if err != nil {
		return nil, err
	}

	return shops, nil
}

func (s *ExternalShopService) SyncProductsByExternalShopId(externalShopId int32) error {
	service, err := s.getEcommerceExternalShopService(constants.SHOPIFY)
	if err != nil {
		return err
	}

	_, err = service.SyncProducts(externalShopId)
	if err != nil {
		return err
	}
	return nil

}
