package service

type IEcommerceService interface {
	GetEcommerceId() int16
	SyncProducts(externalShopId int64) error

	GetExternalProductsByExternalShopId(externalShopId int64, limit int32, offset int32) (interface{}, error)
	CreateProductVariants(externalProductExternalId interface{}) error
}

func ProvideEcommerceServices(shopifyService IShopifyService) map[int16]IEcommerceService {
	var ecommerceServices = make(map[int16]IEcommerceService)

	ecommerceServices[shopifyService.GetEcommerceId()] = shopifyService
	return ecommerceServices
}
