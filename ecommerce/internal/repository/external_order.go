package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalOrderRepository interface {
	IBaseRepository[model.ExternalOrder]
}

type ExternalOrderRepository struct {
	BaseRepository[model.ExternalOrder]
}

func NewExternalOrderRepository(database *database.PostgresqlDatabase) IExternalOrderRepository {
	repo := &ExternalOrderRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalOrderRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalOrder) (*model.ExternalOrder, error) {
	stmt := table.ExternalOrder.INSERT(columnList).MODEL(data).RETURNING(table.ExternalOrder.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalOrderRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExternalOrder) ([]*model.ExternalOrder, error) {
	stmt := table.ExternalOrder.INSERT(columnList).MODELS(data).RETURNING(table.ExternalOrder.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalOrderRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalOrder) (*model.ExternalOrder, error) {
	stmt := table.ExternalOrder.UPDATE(columnList).MODEL(data).WHERE(table.ExternalOrder.IDExternalOrder.EQ(postgres.Int(data.IDExternalOrder))).RETURNING(table.ExternalOrder.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalOrderRepository) GetById(db qrm.Queryable, id int64) (*model.ExternalOrder, error) {
	stmt := table.ExternalOrder.SELECT(table.ExternalOrder.AllColumns).WHERE(table.ExternalOrder.IDExternalOrder.EQ(postgres.Int(int64(id))))

	var data model.ExternalOrder
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
