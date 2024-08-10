package model

import "time"

type CreateProductsVariantsRequest []struct {
	EcommerceID               int16  `json:"ecommerce_id"`
	ExternalProductExternalID string `json:"external_product_external_id"`
}

type GetProductsByShopIdParam struct {
	PaginationParam
}

type GetProductsByShopIdData struct {
	IDProduct    int64                  `json:"id_product" xml:"id_product"`
	Name         string                 `json:"name" xml:"name"`
	Description  *string                `json:"description" xml:"description"`
	Status       string                 `json:"status" xml:"status"`
	Stock        *int32                 `json:"stock" xml:"stock"`
	OptionTitles map[string]interface{} `json:"option_titles" xml:"option_titles"`
	CreatedAt    time.Time              `json:"created_at" xml:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at" xml:"updated_at"`
}
