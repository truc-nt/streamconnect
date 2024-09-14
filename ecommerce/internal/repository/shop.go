package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IShopRepository interface {
	IBaseRepository[model.Shop]
}

type ShopRepository struct {
	BaseRepository[model.Shop]
}

func NewShopRepository(database *database.PostgresqlDatabase) IShopRepository {
	repo := &ShopRepository{}
	repo.Database = database
	return repo
}

func (r *ShopRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Shop) (*model.Shop, error) {
	stmt := table.Shop.INSERT(columnList).MODEL(data).RETURNING(table.Shop.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ShopRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Shop) ([]*model.Shop, error) {
	stmt := table.Shop.INSERT(columnList).MODELS(data).RETURNING(table.Shop.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ShopRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Shop) (*model.Shop, error) {
	stmt := table.Shop.UPDATE(columnList).MODEL(data).WHERE(table.Shop.IDShop.EQ(postgres.Int(data.IDShop))).RETURNING(table.Shop.AllColumns)
	return r.update(db, stmt)
}

func (r *ShopRepository) GetById(db qrm.Queryable, id int64) (*model.Shop, error) {
	stmt := table.Shop.SELECT(table.Shop.AllColumns).WHERE(table.Shop.IDShop.EQ(postgres.Int(int64(id))))

	var data model.Shop
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
