package services

type Services struct {
	Shopify IShopifyService
}

func NewServices(shopify IShopifyService) *Services {
	return &Services{
		Shopify: shopify,
	}
}
