package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalProductShopifyRepository interface {
	IBaseRepository[model.ExternalProductShopify]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExternalProductShopify, error)
	GetByShopifyProductId(db qrm.Queryable, shopifyProductId int64) ([]*model.ExternalProductShopify, error)
	GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExternalProductShopify, error)
	UpdateProductVariant(db qrm.Queryable, productId int64, variantId int64, shopifyVariantId int64) error

	GetExternalProductsByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) (interface{}, error)
}

type ExternalProductShopifyRepository struct {
	BaseRepository[model.ExternalProductShopify]
}

func NewExternalProductShopifyRepository(database *database.PostgresqlDatabase) IExternalProductShopifyRepository {
	repo := &ExternalProductShopifyRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalProductShopifyRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalProductShopify) (*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.INSERT(columnList).MODEL(data).RETURNING(table.ExternalProductShopify.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalProductShopifyRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExternalProductShopify) ([]*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.
		INSERT(columnList).
		MODELS(data).
		RETURNING(table.ExternalProductShopify.AllColumns).
		ON_CONFLICT(
			table.ExternalProductShopify.ShopifyProductID,
			table.ExternalProductShopify.ShopifyVariantID,
		).
		DO_UPDATE(
			postgres.SET(
				table.ExternalProductShopify.Sku.SET(table.ExternalProductShopify.EXCLUDED.Sku),
				table.ExternalProductShopify.Name.SET(table.ExternalProductShopify.EXCLUDED.Name),
				table.ExternalProductShopify.Stock.SET(table.ExternalProductShopify.EXCLUDED.Stock),
				table.ExternalProductShopify.Option.SET(table.ExternalProductShopify.EXCLUDED.Option),
				table.ExternalProductShopify.Price.SET(table.ExternalProductShopify.EXCLUDED.Price),
				table.ExternalProductShopify.ImageURL.SET(table.ExternalProductShopify.EXCLUDED.ImageURL),
			),
		)
	return r.insertMany(db, stmt)
}

func (r *ExternalProductShopifyRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalProductShopify) (*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.UPDATE(columnList).MODEL(data).RETURNING(table.ExternalProductShopify.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalProductShopifyRepository) GetById(db qrm.Queryable, id int64) (*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.SELECT(table.ExternalProductShopify.AllColumns).WHERE(table.ExternalProductShopify.IDExternalProductShopify.EQ(postgres.Int(id)))
	var data model.ExternalProductShopify
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ExternalProductShopifyRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.SELECT(table.ExternalProductShopify.AllColumns).WHERE(table.ExternalProductShopify.FkExternalShop.EQ(postgres.Int(externalShopId)))
	var data []*model.ExternalProductShopify
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalProductShopifyRepository) GetByShopifyProductId(db qrm.Queryable, shopifyProductId int64) ([]*model.ExternalProductShopify, error) {
	stmt := table.ExternalProductShopify.SELECT(table.ExternalProductShopify.AllColumns).WHERE(table.ExternalProductShopify.ShopifyProductID.EQ(postgres.Int(shopifyProductId)))
	var data []*model.ExternalProductShopify
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalProductShopifyRepository) UpdateProductVariant(db qrm.Queryable, productId int64, variantId int64, shopifyVariantId int64) error {
	stmt := table.ExternalProductShopify.UPDATE(
		table.ExternalProductShopify.FkProduct,
		table.ExternalProductShopify.FkVariant,
	).SET(
		table.ExternalProductShopify.FkProduct.SET(postgres.Int(productId)),
		table.ExternalProductShopify.FkVariant.SET(postgres.Int(variantId)),
	).WHERE(
		table.ExternalProductShopify.ShopifyVariantID.EQ(postgres.Int(shopifyVariantId)),
	).RETURNING(table.ExternalProductShopify.AllColumns)

	var data model.ExternalProductShopify
	if err := stmt.Query(db, &data); err != nil {
		return err
	}

	return nil
}

type GetExternalProductsByExternalShopId struct {
	ExternalProductExternalId *int64    `alias:"external_product_shopify.shopify_product_id"`
	ExternalProductName       string    `alias:"external_product_shopify.name"`
	ProductID                 *int64    `alias:"product.id_product"`
	ProductName               string    `alias:"product.name"`
	TotalStock                int32     `alias:"external_product_shopify.stock"`
	MinPrice                  float64   `alias:"external_product_shopify.price"`
	MaxPrice                  float64   `alias:"external_product_shopify.price"`
	ImageUrl                  *string   `alias:"external_product_shopify.image_url"`
	UpdatedAt                 time.Time `alias:"external_product_shopify.updated_at"`
}

func (r *ExternalProductShopifyRepository) GetExternalProductsByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExternalProductShopify.SELECT(
		table.ExternalProductShopify.ShopifyProductID.AS("GetExternalProductsByExternalShopId.ExternalProductExternalId"),
		postgres.Raw(fmt.Sprintf("(array_agg(%s.%s))[1]", "external_product_shopify", table.ExternalProductShopify.Name.Name())).AS("GetExternalProductsByExternalShopId.ExternalProductName"),
		table.Product.IDProduct.AS("GetExternalProductsByExternalShopId.ProductID"),
		table.Product.Name.AS("GetExternalProductsByExternalShopId.ProductName"),
		postgres.SUM(table.ExternalProductShopify.Stock).AS("GetExternalProductsByExternalShopId.TotalStock"),
		postgres.MIN(table.ExternalProductShopify.Price).AS("GetExternalProductsByExternalShopId.MinPrice"),
		postgres.MAX(table.ExternalProductShopify.Price).AS("GetExternalProductsByExternalShopId.MaxPrice"),
		postgres.MAX(table.ExternalProductShopify.UpdatedAt).AS("GetExternalProductsByExternalShopId.UpdatedAt"),
	).FROM(
		table.ExternalProductShopify.
			LEFT_JOIN(table.Product, table.Product.IDProduct.EQ(table.ExternalProductShopify.FkProduct)),
	).GROUP_BY(
		table.ExternalProductShopify.ShopifyProductID,
		table.ExternalProductShopify.Name,
		table.Product.IDProduct,
		table.Product.Name,
	).WHERE(
		table.ExternalProductShopify.FkExternalShop.EQ(postgres.Int(externalShopId)),
	).LIMIT(limit).OFFSET(offset)

	var data []*GetExternalProductsByExternalShopId
	err := stmt.Query(r.GetDefaultDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalProductShopifyRepository) GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExternalProductShopify, error) {
	var postgresExpression []postgres.Expression

	for _, variantId := range variantIds {
		postgresExpression = append(postgresExpression, postgres.Int(variantId))
	}

	stmt := table.ExternalProductShopify.SELECT(table.ExternalProductShopify.AllColumns).WHERE(table.ExternalProductShopify.FkVariant.IN(postgresExpression...))

	var data []*model.ExternalProductShopify
	err := stmt.Query(r.GetDefaultDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
