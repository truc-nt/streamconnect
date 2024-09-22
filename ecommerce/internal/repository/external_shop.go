package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalShopRepository interface {
	IBaseRepository[model.ExtShop]

	GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error)

	CreateOneV1(columnList postgres.ColumnList, data model.ExtShop) func(db qrm.Queryable) (*model.ExtShop, error)
}

type ExternalShopRepository struct {
	BaseRepository[model.ExtShop]
}

func NewExternalShopRepository(database *database.PostgresqlDatabase) IExternalShopRepository {
	repo := &ExternalShopRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalShopRepository) GetById(db qrm.Queryable, id int64) (*model.ExtShop, error) {
	stmt := table.ExtShop.SELECT(
		table.ExtShop.AllColumns,
		table.Ecommerce.Name.AS("ecommerce")).
		FROM(table.ExtShop.
			INNER_JOIN(table.Ecommerce, table.Ecommerce.IDEcommerce.EQ(table.ExtShop.FkEcommerce))).
		WHERE(table.ExtShop.IDExtShop.EQ(postgres.Int(int64(id))))

	var data model.ExtShop
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ExternalShopRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtShop) (*model.ExtShop, error) {
	stmt := table.ExtShop.INSERT(columnList).MODEL(data).RETURNING(table.ExtShop.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalShopRepository) CreateOneV1(columnList postgres.ColumnList, data model.ExtShop) func(db qrm.Queryable) (*model.ExtShop, error) {
	stmt := table.ExtShop.INSERT(columnList).MODEL(data).RETURNING(table.ExtShop.AllColumns)
	return func(db qrm.Queryable) (*model.ExtShop, error) {
		return r.insertOne(db, stmt)
	}
}

func (r *ExternalShopRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExtShop) ([]*model.ExtShop, error) {
	stmt := table.ExtShop.INSERT(columnList).MODELS(data).RETURNING(table.ExtShop.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalShopRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtShop) (*model.ExtShop, error) {
	stmt := table.ExtShop.UPDATE(columnList).MODEL(data).RETURNING(table.ExtShop.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalShopRepository) GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExtShop.SELECT(
		table.ExtShop.AllColumns,
		table.Ecommerce.Name).
		FROM(table.ExtShop.
			INNER_JOIN(table.Ecommerce, table.Ecommerce.IDEcommerce.EQ(table.ExtShop.FkEcommerce))).
		WHERE(table.ExtShop.FkShop.EQ(postgres.Int(int64(shopId)))).LIMIT(int64(limit)).OFFSET(int64(offset))

	var data []*model.ExtShop
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
