package main

import (
	"ecommerce/internal/configs"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
)

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
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseSQLBuilder(template.DefaultSQLBuilder().
						UseTable(func(table metadata.Table) template.TableSQLBuilder {
							if table.Name == "schema_migrations" {
								return template.TableSQLBuilder{
									Skip: true,
								}
							}
							return template.DefaultTableSQLBuilder(table)
						}),
					)
			}).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							if table.Name == "schema_migrations" {
								return template.TableModel{
									Skip: true,
								}
							}
							return template.DefaultTableModel(table).
								UseField(func(columnMetaData metadata.Column) template.TableModelField {
									defaultTableModelField := template.DefaultTableModelField(columnMetaData)
									return defaultTableModelField.UseTags(
										fmt.Sprintf(`json:"%s"`, columnMetaData.Name),
										fmt.Sprintf(`xml:"%s"`, columnMetaData.Name),
									)
								})
						}),
					)
			}),
	)
	if err != nil {
		panic(err)
	}
}
