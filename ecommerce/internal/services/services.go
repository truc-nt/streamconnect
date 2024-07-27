package services

type Services struct {
	ShopifyService          IShopifyService
	ExternalShopService     IExternalShopService
	ExternalShopAuthService IExternalShopAuthService
}

func NewServices(shopifyService IShopifyService, externalShopService IExternalShopService, externalShopAuthService IExternalShopAuthService) *Services {
	return &Services{
		ShopifyService:          shopifyService,
		ExternalShopService:     externalShopService,
		ExternalShopAuthService: externalShopAuthService,
	}
}
