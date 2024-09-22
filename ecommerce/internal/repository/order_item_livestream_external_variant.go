package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderItemLivestreamExternalVariantRepository interface {
	IBaseRepository[model.OrderItemLivestreamExtVariant]
}

type OrderItemLivestreamExternalVariantRepository struct {
	BaseRepository[model.OrderItemLivestreamExtVariant]
}

func NewOrderItemLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) IOrderItemLivestreamExternalVariantRepository {
	repo := &OrderItemLivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *OrderItemLivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItemLivestreamExtVariant) (*model.OrderItemLivestreamExtVariant, error) {
	stmt := table.OrderItemLivestreamExtVariant.INSERT(columnList).MODEL(data).RETURNING(table.OrderItemLivestreamExtVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.OrderItemLivestreamExtVariant) ([]*model.OrderItemLivestreamExtVariant, error) {
	stmt := table.OrderItemLivestreamExtVariant.INSERT(columnList).MODELS(data).RETURNING(table.OrderItemLivestreamExtVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItemLivestreamExtVariant) (*model.OrderItemLivestreamExtVariant, error) {
	stmt := table.OrderItemLivestreamExtVariant.UPDATE(columnList).MODEL(data).WHERE(table.OrderItemLivestreamExtVariant.ID.EQ(postgres.Int(data.ID))).RETURNING(table.OrderItemLivestreamExtVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *OrderItemLivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.OrderItemLivestreamExtVariant, error) {
	stmt := table.OrderItemLivestreamExtVariant.SELECT(table.OrderItemLivestreamExtVariant.AllColumns).WHERE(table.OrderItemLivestreamExtVariant.ID.EQ(postgres.Int(int64(id))))

	var data model.OrderItemLivestreamExtVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
