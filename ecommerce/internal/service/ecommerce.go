package service

import (
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/model"
)

type IEcommerceService interface {
	GetEcommerceId() int16
	SyncVariants(externalShopId int64) error
	GetStockByExternalProductExternalId(externalShopId int64, externalProductIdMappings []string) ([]*model.ExternalVariantStock, error)
	CreateOrder(user *entity.User, address *entity.UserAddress, externalShopId int64, externalOrderItems []*model.ExternalOrderItem, internalDiscount float64) (string, error)

	//GetProductVariantsByExternalProductExternalId(externalProductExternalId interface{}) (interface{}, error)
	//CreateExternalVariants(externalProductExternalId interface{}) error
}

func ProvideEcommerceServices(shopifyService IShopifyService) map[int16]IEcommerceService {
	var ecommerceServices = make(map[int16]IEcommerceService)

	ecommerceServices[shopifyService.GetEcommerceId()] = shopifyService
	return ecommerceServices
}
