package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
	"github.com/samber/lo"
)

type ILivestreamExternalVariantRepository interface {
	IBaseRepository[model.LivestreamExtVariant]

	GetByLivestreamProductId(db qrm.Queryable, livestreamProductId int64) (*GetByLivestreamProductId, error)
	GetByIds(db qrm.Queryable, livestreamExternalVariantIds []int64) ([]*model.LivestreamExtVariant, error)
	CreateManyOnConflict(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamExtVariant) ([]*model.LivestreamExtVariant, error)
	GetVariantById(db qrm.Queryable, id int64) (*GetVariantById, error)
}

type LivestreamExternalVariantRepository struct {
	BaseRepository[model.LivestreamExtVariant]
}

func NewLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) ILivestreamExternalVariantRepository {
	repo := &LivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExtVariant) (*model.LivestreamExtVariant, error) {
	stmt := table.LivestreamExtVariant.INSERT(columnList).MODEL(data).RETURNING(table.LivestreamExtVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamExtVariant) ([]*model.LivestreamExtVariant, error) {
	stmt := table.LivestreamExtVariant.INSERT(columnList).MODELS(data).RETURNING(table.LivestreamExtVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExtVariant) (*model.LivestreamExtVariant, error) {
	stmt := table.LivestreamExtVariant.UPDATE(columnList).MODEL(data).WHERE(table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(postgres.Int(data.IDLivestreamExtVariant))).RETURNING(table.LivestreamExtVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.LivestreamExtVariant, error) {
	stmt := table.LivestreamExtVariant.SELECT(table.LivestreamExtVariant.AllColumns).WHERE(table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(postgres.Int(int64(id))))

	var data model.LivestreamExtVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetByLivestreamProductId struct {
	IDLivestreamProduct int64  `sql:"primary_key" alias:"livestream_product.id_livestream_product" json:"-"`
	IDProduct           int64  `alias:"product.id_product" json:"id_product"`
	Name                string `alias:"product.name" json:"name"`
	Description         string `alias:"product.description" json:"description"`
	ImageURL            string `json:"image_url"`
	LivestreamVariants  []*struct {
		IDLivestreamVariant        int64       `sql:"primary_key" alias:"variant.id_variant" json:"id_variant"`
		Option                     pgtype.JSON `alias:"variant.option" json:"option"`
		LivestreamExternalVariants []*struct {
			IDLivestreamExternalVariant int64   `alias:"livestream_ext_variant.id_livestream_ext_variant" json:"id_livestream_external_variant"`
			IDExternalShop              int64   `alias:"ext_shop.id_ext_shop" json:"-"`
			IDEcommerce                 int16   `alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
			ImageURL                    string  `alias:"image_variant.url" json:"image_url"`
			Quantity                    int64   `alias:"livestream_ext_variant.quantity" json:"quantity"`
			Stock                       int64   `json:"stock"`
			Price                       float64 `alias:"ext_variant.price" json:"price"`
			ExternalProductIdMapping    string  `alias:"ext_variant.ext_product_id_mapping" json:"-"`
			ExternalIdMapping           string  `alias:"ext_variant.ext_id_mapping" json:"-"`
		} `json:"livestream_external_variants"`
	} `json:"livestream_variants"`
}

func (r *LivestreamExternalVariantRepository) GetByLivestreamProductId(db qrm.Queryable, livestreamProductId int64) (*GetByLivestreamProductId, error) {
	innerVariantAlias := table.Variant.AS("inner_variant")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(innerVariantAlias, innerVariantAlias.IDVariant.EQ(table.ImageVariant.FkVariant)),
		).WHERE(table.Variant.IDVariant.EQ(innerVariantAlias.IDVariant)).LIMIT(1)

	stmt := table.LivestreamExtVariant.SELECT(
		table.LivestreamProduct.IDLivestreamProduct,
		table.LivestreamExtVariant.IDLivestreamExtVariant,
		table.ExtShop.FkEcommerce,
		table.LivestreamExtVariant.Quantity,
		table.Variant.Option,
		table.Product.IDProduct,
		table.Product.Name,
		table.Product.Description,
		table.ExtVariant.Price,
		table.ImageVariant.URL,
		table.Variant.IDVariant,
		table.ExtVariant.ExtIDMapping,
		table.ExtVariant.ExtProductIDMapping,
		table.ExtShop.IDExtShop,
		imageUrlSubQuery.AS("GetByLivestreamProductId.ImageURL"),
	).FROM(
		table.LivestreamExtVariant.
			INNER_JOIN(table.LivestreamProduct, table.LivestreamProduct.IDLivestreamProduct.EQ(table.LivestreamExtVariant.FkLivestreamProduct)).
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.LivestreamExtVariant.FkExtVariant)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			LEFT_JOIN(table.ImageVariant, table.ImageVariant.FkVariant.EQ(table.Variant.IDVariant)),
	).WHERE(table.LivestreamExtVariant.FkLivestreamProduct.EQ(postgres.Int(livestreamProductId)))
	var data GetByLivestreamProductId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *LivestreamExternalVariantRepository) GetByIds(db qrm.Queryable, livestreamExternalVariantIds []int64) ([]*model.LivestreamExtVariant, error) {
	ids := lo.Map(livestreamExternalVariantIds, func(id int64, _ int) postgres.Expression {
		return postgres.Int(id)
	})

	stmt := table.LivestreamExtVariant.SELECT(table.LivestreamExtVariant.AllColumns).WHERE(table.LivestreamExtVariant.IDLivestreamExtVariant.IN(ids...))
	data := make([]*model.LivestreamExtVariant, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *LivestreamExternalVariantRepository) CreateManyOnConflict(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamExtVariant) ([]*model.LivestreamExtVariant, error) {
	stmt := table.LivestreamExtVariant.
		INSERT(columnList).
		MODELS(data).
		RETURNING(table.LivestreamExtVariant.AllColumns).
		ON_CONFLICT(
			table.LivestreamExtVariant.FkLivestreamProduct,
			table.LivestreamExtVariant.FkExtVariant,
		).
		DO_UPDATE(
			postgres.SET(
				table.LivestreamExtVariant.Quantity.SET(table.LivestreamExtVariant.EXCLUDED.Quantity),
			),
		)
	return r.insertMany(db, stmt)
}

type GetVariantById struct {
	*model.LivestreamExtVariant
	*model.ExtVariant
	*model.ExtShop
	*model.Variant
	*model.Product
	IDShop int64 `alias:"product.fk_shop" json:"id_shop"`
}

func (r *LivestreamExternalVariantRepository) GetVariantById(db qrm.Queryable, id int64) (*GetVariantById, error) {
	stmt := table.LivestreamExtVariant.SELECT(
		table.LivestreamExtVariant.AllColumns,
		table.ExtVariant.AllColumns,
		table.Variant.AllColumns,
		table.Product.AllColumns,
		table.ExtShop.AllColumns,
	).FROM(
		table.LivestreamExtVariant.
			LEFT_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.LivestreamExtVariant.FkExtVariant)).
			LEFT_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			LEFT_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			LEFT_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)),
	).WHERE(table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(postgres.Int(int64(id))))

	var data GetVariantById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
