//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type OrderAddress struct {
	IDOrderAddress int64  `sql:"primary_key" json:"id_order_address" xml:"id_order_address"`
	FkOrder        int64  `json:"fk_order" xml:"fk_order"`
	Name           string `json:"name" xml:"name"`
	Phone          string `json:"phone" xml:"phone"`
	Address        string `json:"address" xml:"address"`
	City           string `json:"city" xml:"city"`
}
