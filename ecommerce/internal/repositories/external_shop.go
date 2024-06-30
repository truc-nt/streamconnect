package repositories

import (
	"ecommerce/internal/database"

	"ecommerce/internal/models"
	"ecommerce/internal/tables"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalShopRepository interface {
	IBaseRepository[models.ExternalShop]
	GetTable() *tables.ExternalShopTable

	GetById(db qrm.Queryable, id int32) (*models.ExternalShop, error)
	Create(db qrm.Queryable, data models.ExternalShop, columnList postgres.ColumnList) (*models.ExternalShop, error)
	GetByShopId(db qrm.Queryable, shopId int32, limit int32, offset int32) ([]*models.ExternalShop, error)
}

type ExternalShopRepository struct {
	BaseRepository[models.ExternalShop]
	table *tables.ExternalShopTable
}

func NewExternalShopRepository(database *database.PostgresqlDatabase) IExternalShopRepository {
	repo := &ExternalShopRepository{
		table: tables.ExternalShop,
	}
	repo.Database = database
	return repo
}

func (r *ExternalShopRepository) GetTable() *tables.ExternalShopTable {
	return r.table
}

func (r *ExternalShopRepository) GetById(db qrm.Queryable, id int32) (*models.ExternalShop, error) {
	stmt := r.table.SELECT(r.table.AllColumns).WHERE(r.table.IDExternalShop.EQ(postgres.Int(int64(id))))
	return r.queryRow(db, stmt)
}

func (r *ExternalShopRepository) Create(db qrm.Queryable, data models.ExternalShop, columnList postgres.ColumnList) (*models.ExternalShop, error) {
	stmt := r.table.INSERT(columnList).MODEL(data).RETURNING(r.table.AllColumns)
	return r.queryRow(db, stmt)
}

func (r *ExternalShopRepository) GetByShopId(db qrm.Queryable, shopId int32, limit int32, offset int32) ([]*models.ExternalShop, error) {
	stmt := r.table.SELECT(r.table.AllColumns).WHERE(r.table.FkShop.EQ(postgres.Int(int64(shopId)))).LIMIT(int64(limit)).OFFSET(int64(offset))
	return r.queryRows(db, stmt)
}
