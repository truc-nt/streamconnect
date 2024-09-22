package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ILivestreamProductRepository interface {
	IBaseRepository[model.LivestreamProduct]

	FindAllLivestreamId(db qrm.Queryable, livestreamId int64) ([]model.LivestreamProduct, error)
	GetInfoById(db qrm.Queryable, id int64) (*GetInfoById, error)
	GetByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*GetByLivestreamId, error)
	GetByLivestreamIdAndProductId(db qrm.Queryable, livestreamId, productId int64) (*model.LivestreamProduct, error)
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
	stmt := table.LivestreamProduct.SELECT(table.LivestreamProduct.AllColumns).WHERE(table.LivestreamProduct.IDLivestreamProduct.EQ(postgres.Int(int64(id))))

	var data model.LivestreamProduct
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetInfoById struct {
	*model.LivestreamProduct
	Name        string `alias:"product.name" json:"name"`
	Description string `alias:"product.description" json:"description"`
}

func (r *LivestreamProductRepository) GetInfoById(db qrm.Queryable, id int64) (*GetInfoById, error) {
	stmt := table.LivestreamProduct.SELECT(
		table.LivestreamProduct.AllColumns,
		table.Product.Name,
		table.Product.Description,
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
	IDLivestreamProduct int64   `sql:"primary_key" alias:"livestream_product.id_livestream_product" json:"id_livestream_product"`
	Priority            int64   `alias:"livestream_product.priority" json:"priority"`
	Name                string  `alias:"product.name" json:"name"`
	ImageURL            string  `json:"image_url"`
	MinPrice            float64 `json:"min_price"`
	MaxPrice            float64 `json:"max_price"`
}

func (r *LivestreamProductRepository) GetByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*GetByLivestreamId, error) {
	innerProductAlias := table.Product.AS("inner_product")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ImageVariant.FkVariant)).
				LEFT_JOIN(innerProductAlias, innerProductAlias.IDProduct.EQ(table.Variant.FkProduct)),
		).WHERE(table.Product.IDProduct.EQ(innerProductAlias.IDProduct)).LIMIT(1)

	stmt := table.LivestreamProduct.SELECT(
		table.LivestreamProduct.IDLivestreamProduct,
		table.LivestreamProduct.Priority,
		table.Product.Name,
		imageUrlSubQuery.AS("GetByLivestreamId.ImageURL"),
		postgres.MIN(table.ExtVariant.Price).AS("GetByLivestreamId.MinPrice"),
		postgres.MAX(table.ExtVariant.Price).AS("GetByLivestreamId.MaxPrice"),
	).FROM(
		table.LivestreamProduct.
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.LivestreamProduct.FkProduct)).
			INNER_JOIN(table.LivestreamExtVariant, table.LivestreamExtVariant.FkLivestreamProduct.EQ(table.LivestreamProduct.IDLivestreamProduct)).
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.LivestreamExtVariant.FkExtVariant)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			LEFT_JOIN(table.ImageVariant, table.ImageVariant.FkVariant.EQ(table.Variant.IDVariant)),
	).WHERE(
		table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)),
	).GROUP_BY(
		table.Product.IDProduct,
		table.LivestreamProduct.IDLivestreamProduct,
		table.Variant.IDVariant,
	).ORDER_BY(
		table.LivestreamProduct.Priority.ASC(),
	)

	data := make([]*GetByLivestreamId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *LivestreamProductRepository) GetByLivestreamIdAndProductId(db qrm.Queryable, livestreamId, productId int64) (*model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.SELECT(table.LivestreamProduct.AllColumns).WHERE(
		table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)).AND(
			table.LivestreamProduct.FkProduct.EQ(postgres.Int(productId))),
	)

	var data model.LivestreamProduct
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *LivestreamProductRepository) FindAllLivestreamId(db qrm.Queryable, livestreamId int64) ([]model.LivestreamProduct, error) {
	stmt := table.LivestreamProduct.SELECT(table.LivestreamProduct.AllColumns).WHERE(table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)))

	var data []model.LivestreamProduct
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
