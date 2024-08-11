package main

import (
	"ecommerce/internal/configs"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgtype"
)

/*type JsonType map[string]interface{}

func (m *JsonType) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), m)
}

func (m JsonType) Value() (driver.Value, error) {
	return json.Marshal(m)
}*/

var tagsMapping = map[string]interface{}{
	"external_variant": map[string]interface{}{
		"id_variant": `shopify:"ID"`,
		"name":       `shopify:"Title"`,
		"sku":        `shopify:"Sku"`,
		"stock":      `shopify:"InventoryQuantity"`,
		"default":    `shopify:"-"`,
	},
}

func main() {
	c := configs.NewConfig()
	err := postgres.Generate(
		"./tools",
		postgres.DBConnection{
			Host:       c.Postgres.Host,
			Port:       c.Postgres.Port,
			User:       c.Postgres.User,
			Password:   c.Postgres.Password,
			DBName:     c.Postgres.DbName,
			SchemaName: "public",
			SslMode:    "disable",
		},
		/*template.Default(postgres2.Dialect).
		UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
			return template.DefaultSchema(schemaMetaData).
				UseSQLBuilder(template.DefaultSQLBuilder().
					UseTable(func(table metadata.Table) template.TableSQLBuilder {
						if table.Name == "schema_migrations" {
							return template.TableSQLBuilder{
								Skip: true,
							}
						}
						return template.DefaultTableSQLBuilder(table).
							UseColumn(func(column metadata.Column) template.TableSQLBuilderColumn {
								defaultTableSQLBuilderColumn := template.DefaultTableSQLBuilderColumn(column)
								fmt.Println(column.DataType.Name)
								if column.DataType.Name == "json" {
									defaultTableSQLBuilderColumn.Type = "Number"
								}
								return defaultTableSQLBuilderColumn
							})
					}),
				)
		}),*/
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UsePath("../").
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							if table.Name == "schema_migrations" {
								return template.TableModel{
									Skip: true,
								}
							}
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData).UseTags(
										fmt.Sprintf(`json:"%s"`, columnMetaData.Name),
										fmt.Sprintf(`xml:"%s"`, columnMetaData.Name),
									)

									if tags, ok := tagsMapping[table.Name].(map[string]interface{}); ok {
										if tag, ok := tags[columnMetaData.Name].(string); ok {
											defaultTableModelField.Tags = append(defaultTableModelField.Tags, tag)
										} else {
											defaultTableModelField.Tags = append(defaultTableModelField.Tags, tags["default"].(string))
										}
									}

									/*if table.Name == "external_product_shopify" {
										switch columnMetaData.Name {
										case "shopify_product_id":
											defaultTableModelField.Tags = append(defaultTableModelField.Tags, `shopify:"ProductID"`)
											break
										default:
											defaultTableModelField.Tags = append(defaultTableModelField.Tags, `shopify:"-"`)
										}
									}*/

									if columnMetaData.DataType.Name == "json" {
										defaultTableModelField.Type = template.NewType(pgtype.JSON{})
									}
									if columnMetaData.DataType.Name == "jsonb" {
										defaultTableModelField.Type = template.NewType(pgtype.JSONB{})
									}
									return defaultTableModelField
								})
						}),
					).
					UseSQLBuilder(template.DefaultSQLBuilder().
						UseTable(func(table metadata.Table) template.TableSQLBuilder {
							if table.Name == "schema_migrations" {
								return template.TableSQLBuilder{
									Skip: true,
								}
							}
							return template.DefaultTableSQLBuilder(table).
								UseColumn(func(column metadata.Column) template.TableSQLBuilderColumn {
									defaultTableSQLBuilderColumn := template.DefaultTableSQLBuilderColumn(column)
									return defaultTableSQLBuilderColumn
								})
						}),
					)
			}),
	)
	if err != nil {
		panic(err)
	}
}
