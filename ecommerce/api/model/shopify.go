package model

type ShopifyConnectParams struct {
	Hmac string `json:"hmac"`
	Host string `json:"host"`
	Shop string `form:"shop"`
}

type ShopifyRedirectParams struct {
	Code  string `form:"code"`
	Hmac  string `form:"hmac"`
	Host  string `form:"host"`
	Shop  string `form:"shop"`
	State int64  `form:"state"`
}

type ShopifySyncProductsParams struct {
	UserId int32 `form:"user_id"`
}
