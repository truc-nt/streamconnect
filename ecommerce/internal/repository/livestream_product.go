package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type ILivestreamProductRepository interface {
	IBaseRepository[model.LivestreamProduct]

	GetInfoById(db qrm.Queryable, id int64) (*GetInfoById, error)
	GetByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*GetByLivestreamId, error)
}

type LivestreamProductRepository struct {
	BaseRepository[model.LivestreamProduct]
}

func NewLivestreamProductRepository(database *database.PostgresqlDatabase) ILivestreamProductRepository {
	repo := &LivestreamProductRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamProductRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamProduct) (*model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.INSERT(columnList).MODEL(data).RETURNING(table.LivestreamProduct.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamProductRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamProduct) ([]*model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.INSERT(columnList).MODELS(data).RETURNING(table.LivestreamProduct.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamProductRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamProduct) (*model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.UPDATE(columnList).MODEL(data).WHERE(table.LivestreamProduct.IDLivestreamProduct.EQ(postgres.Int(int64(data.IDLivestreamProduct)))).RETURNING(table.LivestreamProduct.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamProductRepository) GetById(db qrm.Queryable, id int64) (*model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.SELECT(table.Livestream.AllColumns).WHERE(table.LivestreamProduct.IDLivestreamProduct.EQ(postgres.Int(int64(id))))

	var data model.LivestreamProduct
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetInfoById struct {
	*model.LivestreamProduct
	Name        string      `alias:"product.name" json:"name"`
	Description string      `alias:"product.description" json:"description"`
	Option      pgtype.JSON `alias:"product.option" json:"option"`
}

func (r *LivestreamProductRepository) GetInfoById(db qrm.Queryable, id int64) (*GetInfoById, error) {
	stmt := table.LivestreamProduct.SELECT(
		table.LivestreamProduct.AllColumns,
		table.Product.Name,
		table.Product.Description,
		table.Product.Option,
	).FROM(
		table.LivestreamProduct.
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.LivestreamProduct.FkProduct)),
	).WHERE(table.LivestreamProduct.IDLivestreamProduct.EQ(postgres.Int(int64(id))))

	var data GetInfoById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetByLivestreamId struct {
	IDLivestreamProduct int64   `alias:"livestream_product.id_livestream_product" json:"id_livestream_product"`
	IDProduct           int64   `alias:"product.id_product" json:"id_product"`
	Name                string  `alias:"product.name" json:"name"`
	MinPrice            float64 `alias:"min_price" json:"min_price"`
	MaxPrice            float64 `alias:"max_price" json:"max_price"`
	Priority            int64   `alias:"livestream_product.priority" json:"priority"`
}

func (r *LivestreamProductRepository) GetByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*GetByLivestreamId, error) {
	stmt := table.LivestreamProduct.SELECT(
		table.LivestreamProduct.IDLivestreamProduct,
		table.Product.IDProduct,
		table.Product.Name,
		postgres.MIN(table.ExternalVariant.Price).AS("GetByLivestreamId.min_price"),
		postgres.MAX(table.ExternalVariant.Price).AS("GetByLivestreamId.max_price"),
		table.LivestreamProduct.Priority,
	).FROM(
		table.LivestreamProduct.
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.LivestreamProduct.FkProduct)).
			INNER_JOIN(table.LivestreamExternalVariant, table.LivestreamExternalVariant.FkLivestreamProduct.EQ(table.LivestreamProduct.IDLivestreamProduct)).
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.LivestreamExternalVariant.FkExternalVariant)),
	).WHERE(
		table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)),
	).GROUP_BY(
		table.Product.IDProduct,
		table.LivestreamProduct.IDLivestreamProduct,
	)

	var data []*GetByLivestreamId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
