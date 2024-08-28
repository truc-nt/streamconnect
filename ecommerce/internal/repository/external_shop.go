package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/table"

	"ecommerce/internal/database/model"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalShopRepository interface {
	IBaseRepository[model.ExternalShop]

	GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error)

	CreateOneV1(columnList postgres.ColumnList, data model.ExternalShop) func(db qrm.Queryable) (*model.ExternalShop, error)
}

type ExternalShopRepository struct {
	BaseRepository[model.ExternalShop]
}

func NewExternalShopRepository(database *database.PostgresqlDatabase) IExternalShopRepository {
	repo := &ExternalShopRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalShopRepository) GetById(db qrm.Queryable, id int64) (*model.ExternalShop, error) {
	stmt := table.ExternalShop.SELECT(
		table.ExternalShop.AllColumns,
		table.Ecommerce.Name.AS("ecommerce")).
		FROM(table.ExternalShop.
			INNER_JOIN(table.Ecommerce, table.Ecommerce.IDEcommerce.EQ(table.ExternalShop.FkEcommerce))).
		WHERE(table.ExternalShop.IDExternalShop.EQ(postgres.Int(int64(id))))

	var data model.ExternalShop
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ExternalShopRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalShop) (*model.ExternalShop, error) {
	stmt := table.ExternalShop.INSERT(columnList).MODEL(data).RETURNING(table.ExternalShop.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalShopRepository) CreateOneV1(columnList postgres.ColumnList, data model.ExternalShop) func(db qrm.Queryable) (*model.ExternalShop, error) {
	stmt := table.ExternalShop.INSERT(columnList).MODEL(data).RETURNING(table.ExternalShop.AllColumns)
	return func(db qrm.Queryable) (*model.ExternalShop, error) {
		return r.insertOne(db, stmt)
	}
}

func (r *ExternalShopRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExternalShop) ([]*model.ExternalShop, error) {
	stmt := table.ExternalShop.INSERT(columnList).MODELS(data).RETURNING(table.ExternalShop.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalShopRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalShop) (*model.ExternalShop, error) {
	stmt := table.ExternalShop.UPDATE(columnList).MODEL(data).RETURNING(table.ExternalShop.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalShopRepository) GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExternalShop.SELECT(
		table.ExternalShop.AllColumns,
		table.Ecommerce.Name).
		FROM(table.ExternalShop.
			INNER_JOIN(table.Ecommerce, table.Ecommerce.IDEcommerce.EQ(table.ExternalShop.FkEcommerce))).
		WHERE(table.ExternalShop.FkShop.EQ(postgres.Int(int64(shopId)))).LIMIT(int64(limit)).OFFSET(int64(offset))

	var data []*model.ExternalShop
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
