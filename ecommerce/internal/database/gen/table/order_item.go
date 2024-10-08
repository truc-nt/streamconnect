//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var OrderItem = newOrderItemTable("public", "order_item", "")

type orderItemTable struct {
	postgres.Table

	// Columns
	IDOrderItem  postgres.ColumnInteger
	FkOrder      postgres.ColumnInteger
	FkExtVariant postgres.ColumnInteger
	Quantity     postgres.ColumnInteger
	UnitPrice    postgres.ColumnFloat
	PaidPrice    postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrderItemTable struct {
	orderItemTable

	EXCLUDED orderItemTable
}

// AS creates new OrderItemTable with assigned alias
func (a OrderItemTable) AS(alias string) *OrderItemTable {
	return newOrderItemTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrderItemTable with assigned schema name
func (a OrderItemTable) FromSchema(schemaName string) *OrderItemTable {
	return newOrderItemTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrderItemTable with assigned table prefix
func (a OrderItemTable) WithPrefix(prefix string) *OrderItemTable {
	return newOrderItemTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrderItemTable with assigned table suffix
func (a OrderItemTable) WithSuffix(suffix string) *OrderItemTable {
	return newOrderItemTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrderItemTable(schemaName, tableName, alias string) *OrderItemTable {
	return &OrderItemTable{
		orderItemTable: newOrderItemTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newOrderItemTableImpl("", "excluded", ""),
	}
}

func newOrderItemTableImpl(schemaName, tableName, alias string) orderItemTable {
	var (
		IDOrderItemColumn  = postgres.IntegerColumn("id_order_item")
		FkOrderColumn      = postgres.IntegerColumn("fk_order")
		FkExtVariantColumn = postgres.IntegerColumn("fk_ext_variant")
		QuantityColumn     = postgres.IntegerColumn("quantity")
		UnitPriceColumn    = postgres.FloatColumn("unit_price")
		PaidPriceColumn    = postgres.FloatColumn("paid_price")
		allColumns         = postgres.ColumnList{IDOrderItemColumn, FkOrderColumn, FkExtVariantColumn, QuantityColumn, UnitPriceColumn, PaidPriceColumn}
		mutableColumns     = postgres.ColumnList{FkOrderColumn, FkExtVariantColumn, QuantityColumn, UnitPriceColumn, PaidPriceColumn}
	)

	return orderItemTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		IDOrderItem:  IDOrderItemColumn,
		FkOrder:      FkOrderColumn,
		FkExtVariant: FkExtVariantColumn,
		Quantity:     QuantityColumn,
		UnitPrice:    UnitPriceColumn,
		PaidPrice:    PaidPriceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
