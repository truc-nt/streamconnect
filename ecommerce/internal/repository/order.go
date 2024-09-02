package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderRepository interface {
	IBaseRepository[model.Order]
}

type OrderRepository struct {
	BaseRepository[model.Order]
}

func NewOrderRepository(database *database.PostgresqlDatabase) IOrderRepository {
	repo := &OrderRepository{}
	repo.Database = database
	return repo
}

func (r *OrderRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Order) (*model.Order, error) {
	stmt := table.Order.INSERT(columnList).MODEL(data).RETURNING(table.Order.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *OrderRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Order) ([]*model.Order, error) {
	stmt := table.Order.INSERT(columnList).MODELS(data).RETURNING(table.Order.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *OrderRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Order) (*model.Order, error) {
	stmt := table.Order.UPDATE(columnList).MODEL(data).WHERE(table.Order.IDOrder.EQ(postgres.Int(data.IDOrder))).RETURNING(table.Order.AllColumns)
	return r.update(db, stmt)
}

func (r *OrderRepository) GetById(db qrm.Queryable, id int64) (*model.Order, error) {
	stmt := table.Order.SELECT(table.Order.AllColumns).WHERE(table.Order.IDOrder.EQ(postgres.Int(int64(id))))

	var data model.Order
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
