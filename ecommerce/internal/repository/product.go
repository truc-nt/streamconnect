package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
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
	IDProduct   int64       `alias:"product.id_product" json:"id_product"`
	Name        string      `alias:"product.name" json:"name"`
	Description string      `alias:"product.description" json:"description"`
	Status      string      `alias:"product.status" json:"status"`
	MinPrice    float64     `alias:"min_price" json:"min_price"`
	MaxPrice    float64     `alias:"max_price" json:"max_price"`
	TotalStock  int32       `alias:"total_stock" json:"total_stock"`
	Option      pgtype.JSON `alias:"product.option_titles" json:"option_titles"`
	ImageUrl    string      `alias:"image_url" json:"image_url"`
	CreatedAt   string      `alias:"product.created_at" json:"created_at"`
	UpdatedAt   string      `alias:"product.updated_at" json:"updated_at"`
}

func (r *ProductRepository) GetByShopId(db qrm.Queryable, shopId int64, limit int64, offset int64) ([]*GetByShopId, error) {
	stmt := table.Product.SELECT(
		table.Product.AllColumns,
		postgres.MIN(table.ExternalProductShopify.Price).AS("GetByShopId.min_price"),
		postgres.MAX(table.ExternalProductShopify.Price).AS("GetByShopId.max_price"),
		postgres.SUM(table.ExternalProductShopify.Stock).AS("GetByShopId.total_stock"),
	).FROM(
		table.Product.
			INNER_JOIN(table.Variant, table.Variant.FkProduct.EQ(table.Product.IDProduct)).
			INNER_JOIN(table.ExternalProductShopify, table.ExternalProductShopify.FkVariant.EQ(table.Variant.IDVariant)),
	).GROUP_BY(
		table.Product.IDProduct,
	).WHERE(table.Product.FkShop.EQ(postgres.Int(shopId))).LIMIT(limit).OFFSET(offset)
	var data []*GetByShopId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
