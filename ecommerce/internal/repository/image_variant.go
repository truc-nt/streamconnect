package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IImageVariantRepository interface {
	IBaseRepository[model.ImageVariant]
}

type ImageVariantRepository struct {
	BaseRepository[model.ImageVariant]
}

func NewImageVariantRepository(database *database.PostgresqlDatabase) IImageVariantRepository {
	repo := &ImageVariantRepository{}
	repo.Database = database
	return repo
}

func (r *ImageVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ImageVariant) (*model.ImageVariant, error) {
	stmt := table.ImageVariant.INSERT(columnList).MODEL(data).RETURNING(table.ImageVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ImageVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ImageVariant) ([]*model.ImageVariant, error) {
	stmt := table.ImageVariant.INSERT(columnList).MODELS(data).RETURNING(table.ImageVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ImageVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ImageVariant) (*model.ImageVariant, error) {
	stmt := table.ImageVariant.UPDATE(columnList).MODEL(data).RETURNING(table.ImageVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *ImageVariantRepository) GetById(db qrm.Queryable, id int64) (*model.ImageVariant, error) {
	stmt := table.ImageVariant.SELECT(table.ImageVariant.AllColumns).WHERE(table.ImageVariant.IDImageVariant.EQ(postgres.Int(id)))
	var data model.ImageVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
