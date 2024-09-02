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
	IBaseRepository[model.LivestreamExternalVariant]

	GetByLivestreamProductId(db qrm.Queryable, livestreamProductId int64) (*GetByLivestreamProductId, error)
	GetByIds(db qrm.Queryable, livestreamExternalVariantIds []int64) ([]*model.LivestreamExternalVariant, error)
}

type LivestreamExternalVariantRepository struct {
	BaseRepository[model.LivestreamExternalVariant]
}

func NewLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) ILivestreamExternalVariantRepository {
	repo := &LivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExternalVariant) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.LivestreamExternalVariant) ([]*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.INSERT(columnList).MODELS(data).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.LivestreamExternalVariant) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(postgres.Int(data.IDLivestreamExternalVariant))).RETURNING(table.LivestreamExternalVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.LivestreamExternalVariant, error) {
	stmt := table.LivestreamExternalVariant.SELECT(table.LivestreamExternalVariant.AllColumns).WHERE(table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(postgres.Int(int64(id))))

	var data model.LivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetByLivestreamProductId struct {
	IDLivestreamProduct int64  `sql:"primary_key" alias:"livestream_product.id_livestream_product" json:"-"`
	Name                string `alias:"product.name" json:"name"`
	Description         string `alias:"product.description" json:"description"`
	ImageURL            string `json:"image_url"`
	LivestreamVariants  []*struct {
		IDLivestreamVariant        int64       `sql:"primary_key" alias:"variant.id_variant" json:"id_variant"`
		Option                     pgtype.JSON `alias:"variant.option" json:"option"`
		LivestreamExternalVariants []*struct {
			IDLivestreamExternalVariant int64   `alias:"livestream_external_variant.id_livestream_external_variant" json:"id_livestream_external_variant"`
			IDEcommerce                 int16   `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
			ImageURL                    string  `alias:"image_variant.url" json:"image_url"`
			Quantity                    int64   `alias:"livestream_external_variant.quantity" json:"quantity"`
			Price                       float64 `alias:"external_variant.price" json:"price"`
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

	stmt := table.LivestreamExternalVariant.SELECT(
		table.LivestreamProduct.IDLivestreamProduct,
		table.LivestreamExternalVariant.IDLivestreamExternalVariant,
		table.ExternalShop.FkEcommerce,
		table.LivestreamExternalVariant.Quantity,
		table.Variant.Option,
		table.Product.Name,
		table.Product.Description,
		table.ExternalVariant.Price,
		table.ImageVariant.URL,
		table.Variant.IDVariant,
		imageUrlSubQuery.AS("GetByLivestreamProductId.ImageURL"),
	).FROM(
		table.LivestreamExternalVariant.
			INNER_JOIN(table.LivestreamProduct, table.LivestreamProduct.IDLivestreamProduct.EQ(table.LivestreamExternalVariant.FkLivestreamProduct)).
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.LivestreamExternalVariant.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			LEFT_JOIN(table.ImageVariant, table.ImageVariant.FkVariant.EQ(table.Variant.IDVariant)),
	).WHERE(table.LivestreamExternalVariant.FkLivestreamProduct.EQ(postgres.Int(livestreamProductId)))
	var data GetByLivestreamProductId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *LivestreamExternalVariantRepository) GetByIds(db qrm.Queryable, livestreamExternalVariantIds []int64) ([]*model.LivestreamExternalVariant, error) {
	ids := lo.Map(livestreamExternalVariantIds, func(id int64, _ int) postgres.Expression {
		return postgres.Int(id)
	})

	stmt := table.LivestreamExternalVariant.SELECT(table.LivestreamExternalVariant.AllColumns).WHERE(table.LivestreamExternalVariant.IDLivestreamExternalVariant.IN(ids...))
	var data []*model.LivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
