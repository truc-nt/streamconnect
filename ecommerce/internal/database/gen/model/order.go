//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Order struct {
	IDOrder          int64     `sql:"primary_key" json:"id_order" xml:"id_order"`
	FkUser           int64     `json:"fk_user" xml:"fk_user"`
	FkShop           int64     `json:"fk_shop" xml:"fk_shop"`
	FkShippingMethod int16     `json:"fk_shipping_method" xml:"fk_shipping_method"`
	FkPaymentMethod  int16     `json:"fk_payment_method" xml:"fk_payment_method"`
	CreatedAt        time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" xml:"updated_at"`
}
