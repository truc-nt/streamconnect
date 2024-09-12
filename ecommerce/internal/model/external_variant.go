package model

import (
	"github.com/jackc/pgtype"
)

type ExternalVariant struct {
	ExternalProductIdMapping *string     `shopify:"-"`
	ExternalIdMapping        string      `shopify:"-"`
	Sku                      *string     `shopify:"Sku"`
	Name                     string      `shopify:"Title"`
	Option                   pgtype.JSON `shopify:"-"`
	Status                   string      `shopify:"-"`
	Price                    float64     `shopify:"-"`
	Url                      string      `shopify:"-"`
	ImageUrl                 string      `shopify:"-"`
}

type ExternalVariantStock struct {
	ExternalProductIdMapping *string `shopify:"-"`
	ExternalIdMapping        string  `shopify:"-"`
	Stock                    int64   `shopify:"InventoryQuantity"`
}

type ExternalOrder struct {
	IDExternalShop      int64   `shopify:"-"`
	IDEcommerce         int16   `shopify:"-"`
	ShippingFee         float64 `shopify:"-"`
	ShippingFeeDiscount float64 `shopify:"-"`
	InternalDiscount    float64 `shopify:"-"`
	ExternalDiscount    float64 `shopify:"-"`
	ExternalOrderItems  []*ExternalOrderItem
}

type ExternalOrderItem struct {
	ExternalIdMapping string `shopify:"-"`
	Quantity          int64  `shopify:"-"`
}
