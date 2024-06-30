package database

import (
	"ecommerce/internal/configs"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	connectionString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
)

type PostgresqlDatabase struct {
	Config *configs.Config
	Db     *sqlx.DB
}

func NewPostgresDatabase(c *configs.Config) *PostgresqlDatabase {
	connStr := getConnectionString(&c.Postgres)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		fmt.Println("Error: Could not establish a connection with the database")
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Error: Could not establish a connection with the database")
		panic(err)
	}

	return &PostgresqlDatabase{
		Config: c,
		Db:     db,
	}
}

func getConnectionString(c *configs.PostgresConfig) string {
	return fmt.Sprintf(connectionString, c.Host, c.User, c.Password, c.DbName, c.Port)
}
