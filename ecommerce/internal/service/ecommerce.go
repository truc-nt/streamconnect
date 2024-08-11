package service

type IEcommerceService interface {
	GetEcommerceId() int16
	SyncProducts(externalShopId int64) error

	GetExternalProductsByExternalShopId(externalShopId int64, limit int64, offset int64) (interface{}, error)
	//GetProductVariantsByExternalProductExternalId(externalProductExternalId interface{}) (interface{}, error)
	CreateExternalVariants(externalProductExternalId interface{}) error
	GetExternalProductByVariantIds(variantIds []int64) (interface{}, error)
}

func ProvideEcommerceServices(shopifyService IShopifyService) map[int16]IEcommerceService {
	var ecommerceServices = make(map[int16]IEcommerceService)

	ecommerceServices[shopifyService.GetEcommerceId()] = shopifyService
	return ecommerceServices
}
