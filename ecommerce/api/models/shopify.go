package models

type ShopifyConnectRequest struct {
	UserId       int32  `json:"user_id"`
	ShopName     string `json:"shop_name"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type ShopifyRedirectParams struct {
	Code         string `form:"code"`
	Shop         string `form:"shop"`
	UserId       int32  `form:"state"`
	ClientId     string `form:"client_id"`
	ClientSecret string `form:"client_secret"`
}

type ShopifySyncProductsParams struct {
	UserId int32 `form:"user_id"`
}
