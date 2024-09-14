package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ICartRepository interface {
	IBaseRepository[model.Cart]
}

type CartRepository struct {
	BaseRepository[model.Cart]
}

func NewCartRepository(database *database.PostgresqlDatabase) ICartRepository {
	repo := &CartRepository{}
	repo.Database = database
	return repo
}

func (r *CartRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Cart) (*model.Cart, error) {
	stmt := table.Cart.INSERT(columnList).MODEL(data).RETURNING(table.Cart.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *CartRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Cart) ([]*model.Cart, error) {
	stmt := table.Cart.INSERT(columnList).MODELS(data).RETURNING(table.Cart.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *CartRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Cart) (*model.Cart, error) {
	stmt := table.Cart.UPDATE(columnList).MODEL(data).WHERE(table.Cart.IDCart.EQ(postgres.Int(data.IDCart))).RETURNING(table.Cart.AllColumns)
	return r.update(db, stmt)
}

func (r *CartRepository) GetById(db qrm.Queryable, id int64) (*model.Cart, error) {
	stmt := table.Cart.SELECT(table.Cart.AllColumns).WHERE(table.Cart.IDCart.EQ(postgres.Int(int64(id))))

	var data model.Cart
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
