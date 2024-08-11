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

type IExternalVariantRepository interface {
	IBaseRepository[model.ExternalVariant]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExternalVariant, error)
	GetByExternalProductId(db qrm.Queryable, externalProductId string) ([]*model.ExternalVariant, error)
	GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExternalVariant, error)
	UpdateExternalVariant(db qrm.Queryable, variantId int64, shopifyVariantId string) error

	GetExternalProductsByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) (interface{}, error)
}

type ExternalVariantRepository struct {
	BaseRepository[model.ExternalVariant]
}

func NewExternalVariantRepository(database *database.PostgresqlDatabase) IExternalVariantRepository {
	repo := &ExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalVariant) (*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.ExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExternalVariant) ([]*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.
		INSERT(columnList).
		MODELS(data).
		RETURNING(table.ExternalVariant.AllColumns).
		ON_CONFLICT(
			table.ExternalVariant.IDExternalProduct,
			table.ExternalVariant.IDExternal,
		).
		DO_UPDATE(
			postgres.SET(
				table.ExternalVariant.Sku.SET(table.ExternalVariant.EXCLUDED.Sku),
				table.ExternalVariant.Name.SET(table.ExternalVariant.EXCLUDED.Name),
				table.ExternalVariant.Stock.SET(table.ExternalVariant.EXCLUDED.Stock),
				table.ExternalVariant.Option.SET(table.ExternalVariant.EXCLUDED.Option),
				table.ExternalVariant.Price.SET(table.ExternalVariant.EXCLUDED.Price),
				table.ExternalVariant.ImageURL.SET(table.ExternalVariant.EXCLUDED.ImageURL),
			),
		)
	return r.insertMany(db, stmt)
}

func (r *ExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalVariant) (*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.UPDATE(columnList).MODEL(data).RETURNING(table.ExternalVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.IDExternalVariant.EQ(postgres.Int(id)))
	var data model.ExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ExternalVariantRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.FkExternalShop.EQ(postgres.Int(externalShopId)))
	var data []*model.ExternalVariant
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetByExternalProductId(db qrm.Queryable, externalProductId string) ([]*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.IDExternalProduct.EQ(postgres.String(externalProductId)))
	var data []*model.ExternalVariant
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) UpdateExternalVariant(db qrm.Queryable, variantId int64, shopifyVariantId string) error {
	stmt := table.ExternalVariant.UPDATE(
		table.ExternalVariant.FkVariant,
	).SET(
		table.ExternalVariant.FkVariant.SET(postgres.Int(variantId)),
	).WHERE(
		table.ExternalVariant.IDExternal.EQ(postgres.String(shopifyVariantId)),
	).RETURNING(table.ExternalVariant.AllColumns)

	var data model.ExternalVariant
	if err := stmt.Query(db, &data); err != nil {
		return err
	}

	return nil
}

type GetExternalProductsByExternalShopId struct {
	ExternalProductExternalId *int64
	ExternalProductName       string
	ProductID                 *int64
	ProductName               string
	TotalStock                int32
	MinPrice                  float64
	MaxPrice                  float64
	ImageUrl                  *string
	UpdatedAt                 time.Time
}

func (r *ExternalVariantRepository) GetExternalProductsByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExternalVariant.SELECT(
		table.ExternalVariant.IDExternalProduct.AS("GetExternalProductsByExternalShopId.ExternalProductExternalId"),
		postgres.Raw(fmt.Sprintf("(array_agg(%s.%s))[1]", "external_variant", table.ExternalVariant.Name.Name())).AS("GetExternalProductsByExternalShopId.ExternalProductName"),
		table.Product.IDProduct.AS("GetExternalProductsByExternalShopId.ProductID"),
		table.Product.Name.AS("GetExternalProductsByExternalShopId.ProductName"),
		postgres.SUM(table.ExternalVariant.Stock).AS("GetExternalProductsByExternalShopId.TotalStock"),
		postgres.MIN(table.ExternalVariant.Price).AS("GetExternalProductsByExternalShopId.MinPrice"),
		postgres.MAX(table.ExternalVariant.Price).AS("GetExternalProductsByExternalShopId.MaxPrice"),
		postgres.MAX(table.ExternalVariant.UpdatedAt).AS("GetExternalProductsByExternalShopId.UpdatedAt"),
	).FROM(
		table.ExternalVariant.
			LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			LEFT_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)),
	).GROUP_BY(
		table.ExternalVariant.IDExternalProduct,
		table.ExternalVariant.Name,
		table.Product.IDProduct,
		table.Product.Name,
	).WHERE(
		table.ExternalVariant.FkExternalShop.EQ(postgres.Int(externalShopId)),
	).LIMIT(limit).OFFSET(offset)

	var data []*GetExternalProductsByExternalShopId
	err := stmt.Query(r.GetDefaultDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExternalVariant, error) {
	var postgresExpression []postgres.Expression

	for _, variantId := range variantIds {
		postgresExpression = append(postgresExpression, postgres.Int(variantId))
	}

	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.FkVariant.IN(postgresExpression...))

	var data []*model.ExternalVariant
	err := stmt.Query(r.GetDefaultDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
