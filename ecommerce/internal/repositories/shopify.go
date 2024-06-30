package repositories

import (
	"ecommerce/internal/database"
	"fmt"

	"ecommerce/internal/models"
	"ecommerce/internal/tables"

	"github.com/go-jet/jet/v2/postgres"
)

type IShopifyRepository interface {
	Create(models.UserShopifyAuth) error
	GetByUserId(userId int) (models.UserShopifyAuth, error)
	UpdateByUserId(userId int, data models.UserShopifyAuth) error
}

type ShopifyRepository struct {
	Database *database.PostgresqlDatabase
}

func NewShopifyRepository(database *database.PostgresqlDatabase) IShopifyRepository {
	return &ShopifyRepository{
		Database: database,
	}
}

func (r *ShopifyRepository) Create(data models.UserShopifyAuth) error {
	stmt := tables.UserShopifyAuth.INSERT(
		tables.UserShopifyAuth.FkUser,
		tables.UserShopifyAuth.ShopName,
		tables.UserShopifyAuth.ClientID,
		tables.UserShopifyAuth.ClientSecret,
	).MODEL(data)

	_, err := stmt.Exec(r.Database.Db)
	fmt.Println(err)
	return err
}

func (r *ShopifyRepository) GetByUserId(userId int) (models.UserShopifyAuth, error) {
	var data models.UserShopifyAuth
	stmt := tables.UserShopifyAuth.SELECT(
		tables.UserShopifyAuth.AllColumns,
	).WHERE(
		tables.UserShopifyAuth.FkUser.EQ(postgres.Int(int64(userId))),
	)

	err := stmt.Query(r.Database.Db, &data)
	return data, err
}

func (r *ShopifyRepository) UpdateByUserId(userId int, data models.UserShopifyAuth) error {
	stmt := tables.UserShopifyAuth.UPDATE(
		tables.UserShopifyAuth.AllColumns,
	).MODEL(data).WHERE(
		tables.UserShopifyAuth.FkUser.EQ(postgres.Int(int64(userId))),
	)

	_, err := stmt.Exec(r.Database.Db)
	return err
}
