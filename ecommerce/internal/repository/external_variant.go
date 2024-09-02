package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IExternalVariantRepository interface {
	IBaseRepository[model.ExternalVariant]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExternalVariant, error)
	GetByExternalProductId(db qrm.Queryable, externalProductId string) ([]*model.ExternalVariant, error)
	GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExternalVariant, error)
	UpdateExternalVariant(db qrm.Queryable, variantId int64, shopifyVariantId string) error

	GetExternalVariantsGroupByProduct(db qrm.Queryable, limit int64, offset int64) (interface{}, error)
	GetExternalVariantsByExternalProductIdMapping(db qrm.Queryable, externalProductIdMapping string) ([]*model.ExternalVariant, error)
	GetExternalVariantInfoById(db qrm.Queryable, id int64) (*GetExternalVariantInfoById, error)
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
			table.ExternalVariant.ExternalProductIDMapping,
			table.ExternalVariant.ExternalIDMapping,
		).
		DO_UPDATE(
			postgres.SET(
				table.ExternalVariant.Sku.SET(table.ExternalVariant.EXCLUDED.Sku),
				table.ExternalVariant.Option.SET(table.ExternalVariant.EXCLUDED.Option),
			),
		)
	return r.insertMany(db, stmt)
}

func (r *ExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExternalVariant) (*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.ExternalVariant.IDExternalVariant.EQ(postgres.Int(data.IDExternalVariant))).RETURNING(table.ExternalVariant.AllColumns)
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
	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.ExternalProductIDMapping.EQ(postgres.String(externalProductId)))
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
		table.ExternalVariant.ExternalIDMapping.EQ(postgres.String(shopifyVariantId)),
	).RETURNING(table.ExternalVariant.AllColumns)

	var data model.ExternalVariant
	if err := stmt.Query(db, &data); err != nil {
		return err
	}

	return nil
}

type GetExternalVariantsGroupByProduct struct {
	ExternalProductIdMapping string  `sql:"primary_key" alias:"external_variant.external_product_id_mapping" json:"external_product_id_mapping"`
	Name                     string  `alias:"external_variant.name" json:"name"`
	ImageUrl                 string  `alias:"external_variant.image_url" json:"image_url"`
	Status                   string  `alias:"external_variant.status" json:"status"`
	IDProduct                *int64  `alias:"product.id_product" json:"id_product"`
	ProductName              *string `alias:"product.name" json:"product_name"`
	ShopName                 string  `alias:"external_shop.name" json:"shop_name"`
	ExternalVariants         []struct {
		IDExternalVariant int64       `alias:"external_variant.id_external_variant" json:"id_external_variant"`
		IDVariant         *int64      `alias:"external_variant.fk_variant" json:"id_variant"`
		Sku               string      `alias:"external_variant.sku" json:"sku"`
		Option            pgtype.JSON `alias:"external_variant.option" json:"option"`
		ImageUrl          string      `alias:"external_variant.image_url" json:"image_url"`
		Price             float64     `alias:"external_variant.price" json:"price"`
	} `json:"external_variants"`
}

func (r *ExternalVariantRepository) GetExternalVariantsGroupByProduct(db qrm.Queryable, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExternalVariant.SELECT(
		table.ExternalVariant.AllColumns,
		table.Product.IDProduct,
		table.Product.Name,
		table.ExternalShop.Name,
	).FROM(
		table.ExternalVariant.
			LEFT_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			LEFT_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)),
	).LIMIT(limit).OFFSET(offset).ORDER_BY(table.ExternalVariant.ExternalProductIDMapping)

	var data []*GetExternalVariantsGroupByProduct
	err := stmt.Query(r.GetDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetExternalVariantsByExternalProductIdMapping(db qrm.Queryable, externalProductIdMapping string) ([]*model.ExternalVariant, error) {
	stmt := table.ExternalVariant.SELECT(table.ExternalVariant.AllColumns).WHERE(table.ExternalVariant.ExternalProductIDMapping.EQ(postgres.String(externalProductIdMapping)))

	var data []*model.ExternalVariant
	err := stmt.Query(r.GetDatabase().Db, &data)
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
	err := stmt.Query(r.GetDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetExternalVariantInfoById struct {
	model.ExternalVariant
	IDShop int64 `alias:"external_shop.fk_shop" json:"id_shop"`
}

func (r *ExternalVariantRepository) GetExternalVariantInfoById(db qrm.Queryable, id int64) (*GetExternalVariantInfoById, error) {
	stmt := table.ExternalVariant.SELECT(
		table.ExternalVariant.AllColumns,
		table.ExternalShop.FkShop,
	).FROM(
		table.ExternalVariant.
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)),
	).WHERE(
		table.ExternalVariant.IDExternalVariant.EQ(postgres.Int(id)),
	)

	var data GetExternalVariantInfoById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
