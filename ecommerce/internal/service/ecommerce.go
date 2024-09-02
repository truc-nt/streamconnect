package service

import "ecommerce/internal/model"

type IEcommerceService interface {
	GetEcommerceId() int16
	SyncVariants(externalShopId int64) error
	GetStockByExternalProductExternalId(externalShopId int64, externalProductIdMappings []string) ([]*model.ExternalVariantStock, error)
	CreateOrder(externalShopId int64, externalOrderItems []*model.ExternalOrderItem) (string, error)

	//GetProductVariantsByExternalProductExternalId(externalProductExternalId interface{}) (interface{}, error)
	//CreateExternalVariants(externalProductExternalId interface{}) error
}

func ProvideEcommerceServices(shopifyService IShopifyService) map[int16]IEcommerceService {
	var ecommerceServices = make(map[int16]IEcommerceService)

	ecommerceServices[shopifyService.GetEcommerceId()] = shopifyService
	return ecommerceServices
}
