package repositories

import (
	"ecommerce/internal/database"

	"ecommerce/internal/models"
	"ecommerce/internal/tables"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IShopifyExternalShopAuthRepository interface {
	IBaseRepository[models.ShopifyExternalShopAuth]
	GetTable() *tables.ShopifyExternalShopAuthTable

	GetByExternalShopId(db qrm.Queryable, externalShopId int32) (*models.ShopifyExternalShopAuth, error)
	Create(db qrm.Queryable, data models.ShopifyExternalShopAuth) (*models.ShopifyExternalShopAuth, error)
}

type ShopifyExternalShopAuthRepository struct {
	BaseRepository[models.ShopifyExternalShopAuth]
	table *tables.ShopifyExternalShopAuthTable
}

func NewShopifyExternalShopAuthRepository(database *database.PostgresqlDatabase) IShopifyExternalShopAuthRepository {
	repo := &ShopifyExternalShopAuthRepository{
		table: tables.ShopifyExternalShopAuth,
	}
	repo.Database = database
	return repo
}

func (r *ShopifyExternalShopAuthRepository) GetTable() *tables.ShopifyExternalShopAuthTable {
	return r.table
}

func (r *ShopifyExternalShopAuthRepository) Create(db qrm.Queryable, data models.ShopifyExternalShopAuth) (*models.ShopifyExternalShopAuth, error) {
	stmt := r.table.INSERT(r.table.FkExternalShop, r.table.Name, r.table.AccessToken).MODEL(data).RETURNING(r.table.AllColumns)

	return r.queryRow(db, stmt)
}

func (r *ShopifyExternalShopAuthRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int32) (*models.ShopifyExternalShopAuth, error) {
	stmt := r.table.SELECT(r.table.AllColumns).WHERE(r.table.FkExternalShop.EQ(postgres.Int(int64(externalShopId))))

	return r.queryRow(db, stmt)
}
