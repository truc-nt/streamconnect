package model

type CreateShopForNewUserRequest struct {
	UserID   int64  `json:"user_id"`
	ShopName string `json:"shop_name"`
}
