package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ILivestreamExternalVariantRepository interface {
	IBaseRepository[model.LivestreamExternalVariant]
}

type LivestreamExternalVariantRepository struct {
	BaseRepository[model.LivestreamExternalVariant]
}

func NewLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) ILivestreamExternalVariantRepository {
	repo := &LivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExternalVariant) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamExternalVariant) ([]*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.INSERT(columnList).MODELS(data).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExternalVariant) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(postgres.Int(data.IDLivestreamExternalVariant))).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.SELECT(table.LivestreamExternalVariant.AllColumns).WHERE(table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(postgres.Int(int64(id))))

	var data model.LivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
