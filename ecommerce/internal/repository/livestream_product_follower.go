package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ILivestreamProductFollowerRepository interface {
	IBaseRepository[model.LivestreamProductFollower]
	FindByProductId(db qrm.Queryable, productId int64) ([]model.LivestreamProductFollower, error)
	GetFollowLivestreamProductsInLivestream(db qrm.Queryable, userId, livestreamId int64) ([]*model.LivestreamProductFollower, error)
	DeleteByLivestreamIdAndUserId(db qrm.Queryable, livestreamProductId, userId int64) error
}

type LivestreamProductFollowerRepository struct {
	BaseRepository[model.LivestreamProductFollower]
}

func NewLivestreamProductFollowerRepository(database *database.PostgresqlDatabase) ILivestreamProductFollowerRepository {
	repo := &LivestreamProductFollowerRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamProductFollowerRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamProductFollower) (*model.LivestreamProductFollower, error) {
	stmt := table.LivestreamProductFollower.INSERT(columnList).MODEL(data).RETURNING(table.LivestreamProductFollower.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamProductFollowerRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamProductFollower) ([]*model.LivestreamProductFollower, error) {
	stmt := table.LivestreamProductFollower.INSERT(columnList).MODELS(data).RETURNING(table.LivestreamProductFollower.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamProductFollowerRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamProductFollower) (*model.LivestreamProductFollower, error) {
	return nil, errors.New("not implemented")
}

func (r *LivestreamProductFollowerRepository) GetById(db qrm.Queryable, id int64) (*model.LivestreamProductFollower, error) {
	return nil, errors.New("not implemented")
}

func (r *LivestreamProductFollowerRepository) FindByProductId(db qrm.Queryable, productId int64) ([]model.LivestreamProductFollower, error) {
	stmt := table.LivestreamProductFollower.SELECT(table.LivestreamProductFollower.AllColumns).WHERE(table.LivestreamProductFollower.FkLivestreamProduct.EQ(postgres.Int(productId)))

	var data []model.LivestreamProductFollower
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *LivestreamProductFollowerRepository) GetFollowLivestreamProductsInLivestream(db qrm.Queryable, userId, livestreamId int64) ([]*model.LivestreamProductFollower, error) {
	stmt := table.LivestreamProductFollower.SELECT(table.LivestreamProductFollower.AllColumns).
		FROM(
			table.LivestreamProductFollower.
				INNER_JOIN(table.LivestreamProduct, table.LivestreamProductFollower.FkLivestreamProduct.EQ(table.LivestreamProduct.IDLivestreamProduct)),
		).WHERE(table.LivestreamProductFollower.FkUser.EQ(postgres.Int(userId))).
		WHERE(table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)))

	var data []*model.LivestreamProductFollower
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *LivestreamProductFollowerRepository) DeleteByLivestreamIdAndUserId(db qrm.Queryable, livestreamProductId, userId int64) error {
	stmt := table.LivestreamProductFollower.DELETE().
		WHERE(table.LivestreamProductFollower.FkUser.EQ(postgres.Int(userId))).
		WHERE(table.LivestreamProductFollower.FkLivestreamProduct.EQ(postgres.Int(livestreamProductId))).
		RETURNING(table.LivestreamProductFollower.AllColumns)

	var data model.LivestreamProductFollower
	err := stmt.Query(db, &data)
	if err != nil {
		return err
	}
	return nil
}
