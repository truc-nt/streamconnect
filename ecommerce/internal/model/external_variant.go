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
	ImageUrl                 string      `shopify:"-"`
}

type ExternalVariantStock struct {
	ExternalProductIdMapping *string `shopify:"-"`
	ExternalIdMapping        string  `shopify:"-"`
	Stock                    int64   `shopify:"InventoryQuantity"`
}

type ExternalOrderItem struct {
	ExternalIdMapping string `shopify:"-"`
	Quantity          int64  `shopify:"-"`
}
