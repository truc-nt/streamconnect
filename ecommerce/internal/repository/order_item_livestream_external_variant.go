package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderItemLivestreamExternalVariantRepository interface {
	IBaseRepository[model.OrderItemLivestreamExternalVariant]
}

type OrderItemLivestreamExternalVariantRepository struct {
	BaseRepository[model.OrderItemLivestreamExternalVariant]
}

func NewOrderItemLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) IOrderItemLivestreamExternalVariantRepository {
	repo := &OrderItemLivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *OrderItemLivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItemLivestreamExternalVariant) (*model.OrderItemLivestreamExternalVariant, error) {
	stmt := table.OrderItemLivestreamExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.OrderItemLivestreamExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.OrderItemLivestreamExternalVariant) ([]*model.OrderItemLivestreamExternalVariant, error) {
	stmt := table.OrderItemLivestreamExternalVariant.INSERT(columnList).MODELS(data).RETURNING(table.OrderItemLivestreamExternalVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItemLivestreamExternalVariant) (*model.OrderItemLivestreamExternalVariant, error) {
	stmt := table.OrderItemLivestreamExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.OrderItemLivestreamExternalVariant.ID.EQ(postgres.Int(data.ID))).RETURNING(table.OrderItemLivestreamExternalVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.OrderItemLivestreamExternalVariant, error) {
	stmt := table.OrderItemLivestreamExternalVariant.SELECT(table.OrderItemLivestreamExternalVariant.AllColumns).WHERE(table.OrderItemLivestreamExternalVariant.ID.EQ(postgres.Int(int64(id))))

	var data model.OrderItemLivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
