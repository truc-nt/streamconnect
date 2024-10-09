package model

import "time"

type GetShopResponse struct {
	IDShop      int64     `json:"id_shop" xml:"id_shop"`
	FkUser      int64     `json:"fk_user" xml:"fk_user"`
	Name        string    `json:"name" xml:"name"`
	Description *string   `json:"description" xml:"description"`
	CreatedAt   time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
	IsFollowing bool      `json:"is_following"`
}

type CreateShopForNewUserRequest struct {
	UserID   int64  `json:"user_id"`
	ShopName string `json:"shop_name"`
}

type UpdateShopRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
