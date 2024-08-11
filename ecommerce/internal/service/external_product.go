package service

import (
	"errors"
)

type IExternalProductService interface {
	GetExternalProductsByExternalShopId(externalShopId int64, limit int64, offset int64) (interface{}, error)
}

type ExternalProductService struct {
	EcommerceService    map[int16]IEcommerceService
	ExternalShopService IExternalShopService
}

func NewExternalProductService(ecommerceService map[int16]IEcommerceService, externalShopService IExternalShopService) IExternalProductService {
	return &ExternalProductService{
		EcommerceService:    ecommerceService,
		ExternalShopService: externalShopService,
	}
}

func (s *ExternalProductService) GetExternalProductsByExternalShopId(externalShopId int64, limit int64, offset int64) (interface{}, error) {
	externalShop, err := s.ExternalShopService.GetExternalShopById(externalShopId)
	if err != nil {
		return nil, err
	}

	ecommerceService, ok := s.EcommerceService[externalShop.FkEcommerce]
	if !ok {
		return nil, errors.New("ecommerce service not found")
	}

	externalProducts, err := ecommerceService.GetExternalProductsByExternalShopId(externalShopId, limit, offset)
	if err != nil {
		return nil, err
	}

	return externalProducts, nil
}
