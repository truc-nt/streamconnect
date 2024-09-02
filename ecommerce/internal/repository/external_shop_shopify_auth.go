package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalShopShopifyAuthRepository interface {
	IBaseRepository[model.ExternalShopShopifyAuth]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64) (*model.ExternalShopShopifyAuth, error)
}

type ExternalShopShopifyAuthRepository struct {
	BaseRepository[model.ExternalShopShopifyAuth]
}

func NewExternalShopShopifyAuthRepository(database *database.PostgresqlDatabase) IExternalShopShopifyAuthRepository {
	repo := &ExternalShopShopifyAuthRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalShopShopifyAuthRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalShopShopifyAuth) (*model.ExternalShopShopifyAuth, error) {
	stmt := table.ExternalShopShopifyAuth.INSERT(columnList).MODEL(data).RETURNING(table.ExternalShopShopifyAuth.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExternalShopShopifyAuth) ([]*model.ExternalShopShopifyAuth, error) {
	stmt := table.ExternalShopShopifyAuth.INSERT(columnList).MODELS(data).RETURNING(table.ExternalShopShopifyAuth.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalShopShopifyAuth) (*model.ExternalShopShopifyAuth, error) {
	stmt := table.ExternalShopShopifyAuth.INSERT(columnList).MODELS(data).RETURNING(table.ExternalShopShopifyAuth.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalShopShopifyAuthRepository) GetById(db qrm.Queryable, id int64) (*model.ExternalShopShopifyAuth, error) {
	stmt := table.ExternalShopShopifyAuth.SELECT(table.ExternalShopShopifyAuth.AllColumns).WHERE(table.ExternalShopShopifyAuth.IDExternalShopShopifyAuth.EQ(postgres.Int(int64(id))))
	var data model.ExternalShopShopifyAuth
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ExternalShopShopifyAuthRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int64) (*model.ExternalShopShopifyAuth, error) {
	stmt := table.ExternalShopShopifyAuth.SELECT(table.ExternalShopShopifyAuth.AllColumns).WHERE(table.ExternalShopShopifyAuth.FkExternalShop.EQ(postgres.Int(externalShopId)))
	var data model.ExternalShopShopifyAuth
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
