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
	IBaseRepository[model.ExtVariant]

	GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExtVariant, error)
	GetByExternalProductId(db qrm.Queryable, externalProductId string) ([]*model.ExtVariant, error)
	GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExtVariant, error)
	UpdateExternalVariant(db qrm.Queryable, variantId int64, shopifyVariantId string) error

	GetExternalVariantsGroupByProduct(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error)
	GetExternalVariantsByExternalProductIdMapping(db qrm.Queryable, externalProductIdMapping string) ([]*model.ExtVariant, error)
	GetExternalVariantInfoById(db qrm.Queryable, id int64) (*GetExternalVariantInfoById, error)
}

type ExternalVariantRepository struct {
	BaseRepository[model.ExtVariant]
}

func NewExternalVariantRepository(database *database.PostgresqlDatabase) IExternalVariantRepository {
	repo := &ExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtVariant) (*model.ExtVariant, error) {
	stmt := table.ExtVariant.INSERT(columnList).MODEL(data).RETURNING(table.ExtVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExtVariant) ([]*model.ExtVariant, error) {
	stmt := table.ExtVariant.
		INSERT(columnList).
		MODELS(data).
		RETURNING(table.ExtVariant.AllColumns).
		ON_CONFLICT(
			table.ExtVariant.ExtProductIDMapping,
			table.ExtVariant.ExtIDMapping,
		).
		DO_UPDATE(
			postgres.SET(
				table.ExtVariant.Sku.SET(table.ExtVariant.EXCLUDED.Sku),
				table.ExtVariant.Option.SET(table.ExtVariant.EXCLUDED.Option),
			),
		)
	return r.insertMany(db, stmt)
}

func (r *ExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtVariant) (*model.ExtVariant, error) {
	stmt := table.ExtVariant.UPDATE(columnList).MODEL(data).WHERE(table.ExtVariant.IDExtVariant.EQ(postgres.Int(data.IDExtVariant))).RETURNING(table.ExtVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.ExtVariant, error) {
	stmt := table.ExtVariant.SELECT(table.ExtVariant.AllColumns).WHERE(table.ExtVariant.IDExtVariant.EQ(postgres.Int(id)))
	var data model.ExtVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ExternalVariantRepository) GetByExternalShopId(db qrm.Queryable, externalShopId int64, limit int64, offset int64) ([]*model.ExtVariant, error) {
	stmt := table.ExtVariant.SELECT(table.ExtVariant.AllColumns).WHERE(table.ExtVariant.FkExtShop.EQ(postgres.Int(externalShopId)))
	data := make([]*model.ExtVariant, 0)
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetByExternalProductId(db qrm.Queryable, externalProductId string) ([]*model.ExtVariant, error) {
	stmt := table.ExtVariant.SELECT(table.ExtVariant.AllColumns).WHERE(table.ExtVariant.ExtProductIDMapping.EQ(postgres.String(externalProductId)))
	data := make([]*model.ExtVariant, 0)
	if err := stmt.Query(db, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) UpdateExternalVariant(db qrm.Queryable, variantId int64, shopifyVariantId string) error {
	stmt := table.ExtVariant.UPDATE(
		table.ExtVariant.FkVariant,
	).SET(
		table.ExtVariant.FkVariant.SET(postgres.Int(variantId)),
	).WHERE(
		table.ExtVariant.ExtIDMapping.EQ(postgres.String(shopifyVariantId)),
	).RETURNING(table.ExtVariant.AllColumns)

	var data model.ExtVariant
	if err := stmt.Query(db, &data); err != nil {
		return err
	}

	return nil
}

type GetExternalVariantsGroupByProduct struct {
	ExternalProductIdMapping string  `sql:"primary_key" alias:"ext_variant.ext_product_id_mapping" json:"external_product_id_mapping"`
	Name                     string  `alias:"ext_variant.name" json:"name"`
	ImageUrl                 string  `alias:"ext_variant.image_url" json:"image_url"`
	Status                   string  `alias:"ext_variant.status" json:"status"`
	IDProduct                *int64  `alias:"product.id_product" json:"id_product"`
	ProductName              *string `alias:"product.name" json:"product_name"`
	ShopName                 string  `alias:"ext_shop.name" json:"shop_name"`
	ExternalVariants         []struct {
		IDExternalVariant int64       `alias:"ext_variant.id_ext_variant" json:"id_external_variant"`
		IDVariant         *int64      `alias:"ext_variant.fk_variant" json:"id_variant"`
		Sku               string      `alias:"ext_variant.sku" json:"sku"`
		Option            pgtype.JSON `alias:"ext_variant.option" json:"option"`
		ImageUrl          string      `alias:"ext_variant.image_url" json:"image_url"`
		Price             float64     `alias:"ext_variant.price" json:"price"`
	} `json:"external_variants"`
}

func (r *ExternalVariantRepository) GetExternalVariantsGroupByProduct(db qrm.Queryable, shopId int64, limit int64, offset int64) (interface{}, error) {
	stmt := table.ExtVariant.SELECT(
		table.ExtVariant.AllColumns,
		table.Product.IDProduct,
		table.Product.Name,
		table.ExtShop.Name,
	).FROM(
		table.ExtVariant.
			LEFT_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			LEFT_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)),
	).WHERE(table.ExtShop.FkShop.EQ(postgres.Int(shopId))).
		LIMIT(limit).OFFSET(offset).ORDER_BY(table.ExtVariant.ExtProductIDMapping)

	data := make([]*GetExternalVariantsGroupByProduct, 0)
	err := stmt.Query(r.GetDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetExternalVariantsByExternalProductIdMapping(db qrm.Queryable, externalProductIdMapping string) ([]*model.ExtVariant, error) {
	stmt := table.ExtVariant.SELECT(table.ExtVariant.AllColumns).WHERE(table.ExtVariant.ExtProductIDMapping.EQ(postgres.String(externalProductIdMapping)))

	data := make([]*model.ExtVariant, 0)
	err := stmt.Query(r.GetDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *ExternalVariantRepository) GetByVariantIds(db qrm.Queryable, variantIds []int64) ([]*model.ExtVariant, error) {
	var postgresExpression []postgres.Expression

	for _, variantId := range variantIds {
		postgresExpression = append(postgresExpression, postgres.Int(variantId))
	}

	stmt := table.ExtVariant.SELECT(table.ExtVariant.AllColumns).WHERE(table.ExtVariant.FkVariant.IN(postgresExpression...))

	data := make([]*model.ExtVariant, 0)
	err := stmt.Query(r.GetDatabase().Db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetExternalVariantInfoById struct {
	model.ExtVariant
	IDShop int64 `alias:"ext_shop.fk_shop" json:"id_shop"`
}

func (r *ExternalVariantRepository) GetExternalVariantInfoById(db qrm.Queryable, id int64) (*GetExternalVariantInfoById, error) {
	stmt := table.ExtVariant.SELECT(
		table.ExtVariant.AllColumns,
		table.ExtShop.FkShop,
	).FROM(
		table.ExtVariant.
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)),
	).WHERE(
		table.ExtVariant.IDExtVariant.EQ(postgres.Int(id)),
	)

	var data GetExternalVariantInfoById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
