package service

import (
	"ecommerce/internal/constants"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/table"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
)

type IExternalShopService interface {
	CreateExternalShopShopify(name string) (int64, error)
	GetExternalShopById(externalShopId int64) (*model.ExternalShop, error)
	GetExternalShopsByShopId(shopId int64, limit int32, offset int32) (interface{}, error)
	SyncExternalProductsByExternalShopId(externalShopId int64) error
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

func (s *ExternalShopService) CreateExternalShopShopify(name string) (int64, error) {
	newExternalShop := model.ExternalShop{
		Name:        name,
		FkShop:      1,
		FkEcommerce: constants.SHOPIFY,
	}

	newData, err := s.Repository.CreateOne(
		s.Repository.GetDefaultDatabase().Db,
		postgres.ColumnList{table.ExternalShop.Name, table.ExternalShop.FkEcommerce},
		newExternalShop,
	)
	if err != nil {
		return 0, err
	}

	return newData.IDExternalShop, nil
}

func (s *ExternalShopService) GetExternalShopById(externalShopId int64) (*model.ExternalShop, error) {
	externalShop, err := s.Repository.GetById(s.Repository.GetDefaultDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	return externalShop, nil
}

func (s *ExternalShopService) GetExternalShopsByShopId(shopId int64, limit int32, offset int32) (interface{}, error) {
	externalShops, err := s.Repository.GetByShopId(s.Repository.GetDefaultDatabase().Db, shopId, limit, offset)
	if err != nil {
		return nil, err
	}

	return externalShops, nil
}

func (s *ExternalShopService) SyncExternalProductsByExternalShopId(externalShopId int64) error {
	externalShop, err := s.Repository.GetById(s.Repository.GetDefaultDatabase().Db, externalShopId)
	if err != nil {
		return err
	}

	service, ok := s.EcommerceService[externalShop.FkEcommerce]
	if !ok {
		return fmt.Errorf("ecommerce service not found")
	}

	if err := service.SyncProducts(externalShopId); err != nil {
		return err
	}
	return nil

}
