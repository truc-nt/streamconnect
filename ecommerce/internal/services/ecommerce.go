package services

type IEcommerceService interface {
	GetEcommerceId() int32
	SyncProducts(externalShopId int32) (interface{}, error)
}

func ProvideEcommerceServices(shopifyService IShopifyService) []IEcommerceService {
	return []IEcommerceService{
		shopifyService,
	}
}
