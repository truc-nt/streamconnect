package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/model"
	"ecommerce/internal/database/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ILivestreamRepository interface {
	IBaseRepository[model.Livestream]
}

type LivestreamRepository struct {
	BaseRepository[model.Livestream]
}

func NewLivestreamRepository(database *database.PostgresqlDatabase) ILivestreamRepository {
	repo := &LivestreamRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Livestream) (*model.Livestream, error) {
	stmt := table.Livestream.INSERT(columnList).MODEL(data).RETURNING(table.Livestream.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Livestream) ([]*model.Livestream, error) {
	stmt := table.Livestream.INSERT(columnList).MODELS(data).RETURNING(table.Livestream.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Livestream) (*model.Livestream, error) {
	stmt := table.Livestream.UPDATE(columnList).MODEL(data).WHERE(table.Livestream.IDLivestream.EQ(postgres.Int(data.IDLivestream))).RETURNING(table.Livestream.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamRepository) GetById(db qrm.Queryable, id int64) (*model.Livestream, error) {
	stmt := table.Livestream.SELECT(table.Livestream.AllColumns).WHERE(table.Livestream.IDLivestream.EQ(postgres.Int(int64(id))))

	var data model.Livestream
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
