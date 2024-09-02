package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IProductRepository interface {
	IBaseRepository[model.Product]

	GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) ([]*GetByShopId, error)
}

type ProductRepository struct {
	BaseRepository[model.Product]
}

func NewProductRepository(database *database.PostgresqlDatabase) IProductRepository {
	repo := &ProductRepository{}
	repo.Database = database
	return repo
}

func (r *ProductRepository) GetById(db qrm.Queryable, id int64) (*model.Product, error) {
	stmt := table.Product.SELECT(table.Product.AllColumns).WHERE(table.Product.IDProduct.EQ(postgres.Int(int64(id))))

	var data model.Product
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ProductRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Product) (*model.Product, error) {
	stmt := table.Product.INSERT(columnList).MODEL(data).RETURNING(table.Product.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ProductRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Product) ([]*model.Product, error) {
	stmt := table.Product.INSERT(columnList).MODELS(data).RETURNING(table.Product.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ProductRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Product) (*model.Product, error) {
	stmt := table.Product.UPDATE(columnList).MODEL(data).WHERE(table.Product.IDProduct.EQ(postgres.Int(data.IDProduct))).RETURNING(table.Product.AllColumns)
	return r.update(db, stmt)
}

type GetByShopId struct {
	model.Product
	ImageURL string `json:"image_url"`
}

func (r *ProductRepository) GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) ([]*GetByShopId, error) {
	innerProductAlias := table.Product.AS("inner_product")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ImageVariant.FkVariant)).
				LEFT_JOIN(innerProductAlias, innerProductAlias.IDProduct.EQ(table.Variant.FkProduct)),
		).WHERE(table.Product.IDProduct.EQ(innerProductAlias.IDProduct)).LIMIT(1)

	stmt := table.Product.SELECT(
		table.Product.AllColumns,
		imageUrlSubQuery.AS("GetByShopId.ImageURL"),
	).FROM(
		table.Product.
			INNER_JOIN(table.Variant, table.Variant.FkProduct.EQ(table.Product.IDProduct)),
	).GROUP_BY(
		table.Product.IDProduct,
	).WHERE(table.Product.FkShop.EQ(postgres.Int(shopId))).LIMIT(limit).OFFSET(offset)
	data := make([]*GetByShopId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
