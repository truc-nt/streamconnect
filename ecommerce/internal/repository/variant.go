package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/model"
	"ecommerce/internal/database/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IVariantRepository interface {
	IBaseRepository[model.Variant]

	GetVariantsByProductId(db qrm.Queryable, shopId int64, limit int64, offset int64) ([]*GetVariantsByProductId, error)
}

type VariantRepository struct {
	BaseRepository[model.Variant]
}

func NewVariantRepository(database *database.PostgresqlDatabase) IVariantRepository {
	repo := &VariantRepository{}
	repo.Database = database
	return repo
}

func (r *VariantRepository) GetById(db qrm.Queryable, id int64) (*model.Variant, error) {
	stmt := table.Variant.SELECT(table.Variant.AllColumns).WHERE(table.Variant.IDVariant.EQ(postgres.Int(int64(id))))

	var data model.Variant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *VariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Variant) (*model.Variant, error) {
	stmt := table.Variant.INSERT(columnList).MODEL(data).RETURNING(table.Variant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *VariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Variant) ([]*model.Variant, error) {
	stmt := table.Variant.INSERT(columnList).MODELS(data).RETURNING(table.Variant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *VariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Variant) (*model.Variant, error) {
	stmt := table.Variant.UPDATE(columnList).MODEL(data).RETURNING(table.Variant.AllColumns)
	return r.update(db, stmt)
}

type GetVariantsByProductId struct {
	IDVariant        int64       `sql:"primary_key" alias:"variant.id_variant" json:"id_variant"`
	Sku              string      `alias:"variant.sku" json:"sku"`
	Option           pgtype.JSON `alias:"variant.option" json:"option"`
	Status           string      `alias:"variant.status" json:"status"`
	ExternalVariants []*struct {
		IDExternalVariant        int64   `alias:"external_variant.id_external_variant" json:"id_external_variant"`
		ExternalProductIdMapping string  `alias:"external_variant.external_product_id_mapping" json:"-"`
		ExternalIdMapping        string  `alias:"external_variant.external_id_mapping" json:"-"`
		IDEcommerce              int16   `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
		IDExternalShop           int64   `alias:"external_variant.fk_external_shop"`
		Sku                      string  `alias:"external_variant.sku" json:"sku"`
		Price                    float64 `alias:"external_variant.price" json:"price"`
		Stock                    int64   `json:"stock"`
	} `json:"external_variants"`
}

func (r *VariantRepository) GetVariantsByProductId(db qrm.Queryable, productId int64, limit int64, offset int64) ([]*GetVariantsByProductId, error) {
	stmt := table.Variant.SELECT(
		table.Variant.AllColumns,
		table.ExternalVariant.AllColumns,
		table.ExternalShop.AllColumns,
	).FROM(
		table.Variant.
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.FkVariant.EQ(table.Variant.IDVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)),
	).WHERE(
		table.Variant.FkProduct.EQ(postgres.Int(productId)),
	).LIMIT(limit).OFFSET(offset)

	data := make([]*GetVariantsByProductId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
