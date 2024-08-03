package model

type CreateProductsVariantsRequest []struct {
	EcommerceID               int16  `json:"ecommerce_id"`
	ExternalProductExternalID string `json:"external_product_external_id"`
}
