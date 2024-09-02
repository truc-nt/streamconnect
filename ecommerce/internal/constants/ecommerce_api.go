package constants

const (
	ShopifyRedirectPath           = "/api/shopify/redirect"
	ShopifyBaseURL                = "https://%s.myshopify.com"
	ShopifyTokenKey               = "X-Shopify-Access-Token"
	ShopifyAuthorizePath          = "/admin/oauth/authorize"
	ShopifyAccessTokenPath        = "/admin/oauth/access_token"
	ShopifyGetProductsPath        = "/admin/api/2024-07/products.json"
	ShopifyGetProductVariantsPath = "/admin/api/2024-07/products/%d/variants.json"
	ShopifyCreateOrderPath        = "/admin/api/2024-07/orders.json"
)
