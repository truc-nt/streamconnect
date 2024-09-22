package service

import (
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/repository"
	"fmt"
)

type IExternalShopService interface {
	GetExternalShopById(externalShopId int64) (*model.ExtShop, error)
	GetExternalShopsByShopId(shopId int64, limit int64, offset int64) (interface{}, error)
	SyncExternalVariantsByExternalShopId(externalShopId int64) error
}

type ExternalShopService struct {
	Repository repository.IExternalShopRepository

	EcommerceService map[int16]IEcommerceService
}

func NewExternalShopService(repo repository.IExternalShopRepository, ecommerceService map[int16]IEcommerceService) IExternalShopService {
	return &ExternalShopService{
		Repository:       repo,
		EcommerceService: ecommerceService,
	}
}

func (s *ExternalShopService) GetExternalShopById(externalShopId int64) (*model.ExtShop, error) {
	externalShop, err := s.Repository.GetById(s.Repository.GetDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	return externalShop, nil
}

func (s *ExternalShopService) GetExternalShopsByShopId(shopId int64, limit int64, offset int64) (interface{}, error) {
	externalShops, err := s.Repository.GetByShopId(s.Repository.GetDatabase().Db, shopId, limit, offset)
	if err != nil {
		return nil, err
	}

	return externalShops, nil
}

func (s *ExternalShopService) SyncExternalVariantsByExternalShopId(externalShopId int64) error {
	externalShop, err := s.Repository.GetById(s.Repository.GetDatabase().Db, externalShopId)
	if err != nil {
		return err
	}

	service, ok := s.EcommerceService[externalShop.FkEcommerce]
	if !ok {
		return fmt.Errorf("ecommerce service not found")
	}

	if err := service.SyncVariants(externalShopId); err != nil {
		return err
	}
	return nil

}
