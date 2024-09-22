package model

type CreateOrderWithCartItemsRequest struct {
	IDUserAddress  int64 `json:"id_user_address"`
	IDShop         int64 `json:"id_shop"`
	ExternalOrders []*struct {
		CartItemIds         []int64 `json:"cart_item_ids"`
		ShippingFee         float64 `json:"shipping_fee"`
		ShippingFeeDiscount float64 `json:"shipping_fee_discount"`
		ExternalDiscount    float64 `json:"external_discount"`
		VoucherIds          []int64 `json:"voucher_ids"`
	} `json:"external_orders"`
}
