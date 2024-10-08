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

var Notification = newNotificationTable("public", "notification", "")

type notificationTable struct {
	postgres.Table

	// Columns
	IDNotification postgres.ColumnInteger
	FkUser         postgres.ColumnInteger
	Title          postgres.ColumnString
	Message        postgres.ColumnString
	Type           postgres.ColumnString
	Status         postgres.ColumnString
	RedirectURL    postgres.ColumnString
	CreatedAt      postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NotificationTable struct {
	notificationTable

	EXCLUDED notificationTable
}

// AS creates new NotificationTable with assigned alias
func (a NotificationTable) AS(alias string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new NotificationTable with assigned schema name
func (a NotificationTable) FromSchema(schemaName string) *NotificationTable {
	return newNotificationTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new NotificationTable with assigned table prefix
func (a NotificationTable) WithPrefix(prefix string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new NotificationTable with assigned table suffix
func (a NotificationTable) WithSuffix(suffix string) *NotificationTable {
	return newNotificationTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newNotificationTable(schemaName, tableName, alias string) *NotificationTable {
	return &NotificationTable{
		notificationTable: newNotificationTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newNotificationTableImpl("", "excluded", ""),
	}
}

func newNotificationTableImpl(schemaName, tableName, alias string) notificationTable {
	var (
		IDNotificationColumn = postgres.IntegerColumn("id_notification")
		FkUserColumn         = postgres.IntegerColumn("fk_user")
		TitleColumn          = postgres.StringColumn("title")
		MessageColumn        = postgres.StringColumn("message")
		TypeColumn           = postgres.StringColumn("type")
		StatusColumn         = postgres.StringColumn("status")
		RedirectURLColumn    = postgres.StringColumn("redirect_url")
		CreatedAtColumn      = postgres.TimestampColumn("created_at")
		allColumns           = postgres.ColumnList{IDNotificationColumn, FkUserColumn, TitleColumn, MessageColumn, TypeColumn, StatusColumn, RedirectURLColumn, CreatedAtColumn}
		mutableColumns       = postgres.ColumnList{FkUserColumn, TitleColumn, MessageColumn, TypeColumn, StatusColumn, RedirectURLColumn, CreatedAtColumn}
	)

	return notificationTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		IDNotification: IDNotificationColumn,
		FkUser:         FkUserColumn,
		Title:          TitleColumn,
		Message:        MessageColumn,
		Type:           TypeColumn,
		Status:         StatusColumn,
		RedirectURL:    RedirectURLColumn,
		CreatedAt:      CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
