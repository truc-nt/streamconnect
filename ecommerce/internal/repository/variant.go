package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IVariantRepository interface {
	IBaseRepository[model.Variant]
}

type VariantRepository struct {
	BaseRepository[model.Variant]
}

func NewVariantRepository(database *database.PostgresqlDatabase) IVariantRepository {
	repo := &VariantRepository{}
	repo.Database = database
	return repo
}

func (r *VariantRepository) GetById(db qrm.Queryable, id int64) (*model.Variant, error) {
	stmt := table.Variant.SELECT(table.Variant.AllColumns).WHERE(table.Variant.IDVariant.EQ(postgres.Int(int64(id))))

	var data model.Variant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *VariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Variant) (*model.Variant, error) {
	stmt := table.Variant.INSERT(columnList).MODEL(data).RETURNING(table.Variant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *VariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Variant) ([]*model.Variant, error) {
	stmt := table.Variant.INSERT(columnList).MODELS(data).RETURNING(table.Variant.AllColumns)
	return r.insertMany(db, stmt)
}
