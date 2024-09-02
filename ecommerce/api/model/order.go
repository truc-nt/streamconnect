package model

type CreateOrderWithCartItemsRequest struct {
	IDUser      int64   `json:"id_user"`
	IDCartItems []int64 `json:"id_cart_items"`
	Address     string  `json:"address"`
}
