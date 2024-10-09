package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IProductRepository interface {
	IBaseRepository[model.Product]

	GetProductInfoById(db qrm.Queryable, id int64) (*GetProductInfoById, error)
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

type GetProductInfoById struct {
	IDProduct   int64  `sql:"primary_key" alias:"product.id_product" json:"id_product"`
	Name        string `alias:"product.name" json:"name"`
	Description string `alias:"product.description" json:"description"`
	ImageURL    string `json:"image_url"`
	Variants    []*struct {
		IDVariant        int64       `sql:"primary_key" alias:"variant.id_variant" json:"id_variant"`
		Sku              string      `alias:"variant.sku" json:"sku"`
		Status           string      `alias:"variant.status" json:"status"`
		Option           pgtype.JSON `alias:"variant.option" json:"option"`
		ExternalVariants []*struct {
			IDExtVariant             int64   `sql:"primary_key" alias:"ext_variant.id_ext_variant" json:"id_ext_variant"`
			Price                    float64 `alias:"ext_variant.price" json:"price"`
			ExternalProductIdMapping string  `alias:"ext_variant.ext_product_id_mapping" json:"-"`
			ExternalIdMapping        string  `alias:"ext_variant.ext_id_mapping" json:"-"`
			Stock                    int64   `alias:"ext_variant.stock" json:"stock"`
			ImageURL                 string  `alias:"ext_variant.image_url" json:"image_url"`
			IDEcommerce              int16   `alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
			IDExternalShop           int64   `alias:"ext_shop.id_ext_shop" json:"id_external_shop"`
			ShopName                 string  `alias:"ext_shop.name" json:"shop_name"`
		} `json:"external_variants"`
	} `json:"variants"`
}

func (r *ProductRepository) GetProductInfoById(db qrm.Queryable, id int64) (*GetProductInfoById, error) {
	innerProductAlias := table.Product.AS("inner_product")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ImageVariant.FkVariant)).
				LEFT_JOIN(innerProductAlias, innerProductAlias.IDProduct.EQ(table.Variant.FkProduct)),
		).WHERE(table.Product.IDProduct.EQ(innerProductAlias.IDProduct)).LIMIT(1)

	stmt := table.Product.SELECT(
		table.Product.IDProduct,
		table.Product.Name,
		table.Product.Description,
		table.Variant.IDVariant,
		table.Variant.Sku,
		table.Variant.Status,
		table.Variant.Option,
		table.ExtVariant.IDExtVariant,
		table.ExtVariant.Price,
		table.ExtVariant.ExtProductIDMapping,
		table.ExtVariant.ExtIDMapping,
		table.ExtVariant.ImageURL,
		table.ImageVariant.URL,
		table.ExtShop.IDExtShop,
		table.ExtShop.FkEcommerce,
		table.ExtShop.Name,
		imageUrlSubQuery.AS("GetProductInfoById.ImageURL"),
	).FROM(
		table.Product.
			INNER_JOIN(table.Variant, table.Variant.FkProduct.EQ(table.Product.IDProduct)).
			INNER_JOIN(table.ExtVariant, table.ExtVariant.FkVariant.EQ(table.Variant.IDVariant)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			LEFT_JOIN(table.ImageVariant, table.ImageVariant.FkVariant.EQ(table.Variant.IDVariant)),
	).WHERE(table.Product.IDProduct.EQ(postgres.Int(id)))
	var data GetProductInfoById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
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
