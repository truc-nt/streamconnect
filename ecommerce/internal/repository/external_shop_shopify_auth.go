package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalShopShopifyAuthRepository interface {
	IBaseRepository[model.ExtShopShopifyAuth]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64) (*model.ExtShopShopifyAuth, error)
}

type ExternalShopShopifyAuthRepository struct {
	BaseRepository[model.ExtShopShopifyAuth]
}

func NewExternalShopShopifyAuthRepository(database *database.PostgresqlDatabase) IExternalShopShopifyAuthRepository {
	repo := &ExternalShopShopifyAuthRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalShopShopifyAuthRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtShopShopifyAuth) (*model.ExtShopShopifyAuth, error) {
	stmt := table.ExtShopShopifyAuth.INSERT(columnList).MODEL(data).RETURNING(table.ExtShopShopifyAuth.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExtShopShopifyAuth) ([]*model.ExtShopShopifyAuth, error) {
	stmt := table.ExtShopShopifyAuth.INSERT(columnList).MODELS(data).RETURNING(table.ExtShopShopifyAuth.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtShopShopifyAuth) (*model.ExtShopShopifyAuth, error) {
	stmt := table.ExtShopShopifyAuth.INSERT(columnList).MODELS(data).RETURNING(table.ExtShopShopifyAuth.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) GetById(db qrm.Queryable, id int64) (*model.ExtShopShopifyAuth, error) {
	stmt := table.ExtShopShopifyAuth.SELECT(table.ExtShopShopifyAuth.AllColumns).WHERE(table.ExtShopShopifyAuth.IDExtShopShopifyAuth.EQ(postgres.Int(int64(id))))
	var data model.ExtShopShopifyAuth
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ExternalShopShopifyAuthRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int64) (*model.ExtShopShopifyAuth, error) {
	stmt := table.ExtShopShopifyAuth.SELECT(table.ExtShopShopifyAuth.AllColumns).WHERE(table.ExtShopShopifyAuth.FkExtShop.EQ(postgres.Int(externalShopId)))
	var data model.ExtShopShopifyAuth
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
