package model

import "time"

type GetExternalProductsByExternalShopIdParam struct {
	PaginationParam
}

type GetExternalProductsByExternalShopIdData struct {
	ExternalProductExternalId int64     `json:"external_product_external_id"`
	ExternalProductName       string    `json:"external_product_name"`
	ProductID                 *int64    `json:"product_id"`
	ProductName               string    `json:"product_name"`
	TotalStock                int32     `json:"total_stock"`
	MinPrice                  float64   `json:"min_price"`
	MaxPrice                  float64   `json:"max_price"`
	ImageUrl                  *string   `json:"image_url"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

type GetExternalProductsByShopIdParam struct {
	PaginationParam
}
