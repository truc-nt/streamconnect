package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderItemRepository interface {
	IBaseRepository[model.OrderItem]
}

type OrderItemRepository struct {
	BaseRepository[model.OrderItem]
}

func NewOrderItemRepository(database *database.PostgresqlDatabase) IOrderItemRepository {
	repo := &OrderItemRepository{}
	repo.Database = database
	return repo
}

func (r *OrderItemRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItem) (*model.OrderItem, error) {
	stmt := table.OrderItem.INSERT(columnList).MODEL(data).RETURNING(table.OrderItem.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *OrderItemRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.OrderItem) ([]*model.OrderItem, error) {
	stmt := table.OrderItem.INSERT(columnList).MODELS(data).RETURNING(table.OrderItem.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *OrderItemRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItem) (*model.OrderItem, error) {
	stmt := table.OrderItem.UPDATE(columnList).MODEL(data).WHERE(table.OrderItem.IDOrderItem.EQ(postgres.Int(data.IDOrderItem))).RETURNING(table.OrderItem.AllColumns)
	return r.update(db, stmt)
}

func (r *OrderItemRepository) GetById(db qrm.Queryable, id int64) (*model.OrderItem, error) {
	stmt := table.OrderItem.SELECT(table.OrderItem.AllColumns).WHERE(table.OrderItem.IDOrderItem.EQ(postgres.Int(int64(id))))

	var data model.OrderItem
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
