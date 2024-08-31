package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type ILivestreamExternalVariantRepository interface {
	IBaseRepository[model.LivestreamExternalVariant]

	GetByLivestreamProductId(db qrm.Queryable, livestreamProductId int64) ([]*GetByLivestreamProductId, error)
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
	IDLivestreamExternalVariant int64       `alias:"livestream_external_variant.id_livestream_external_variant" json:"id_livestream_external_variant"`
	IDEcommerce                 int16       `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
	Quantity                    int64       `alias:"livestream_external_variant.quantity" json:"quantity"`
	Option                      pgtype.JSON `alias:"variant.option" json:"option"`
	Price                       float64     `alias:"external_variant.price" json:"price"`
}

func (r *LivestreamExternalVariantRepository) GetByLivestreamProductId(db qrm.Queryable, livestreamProductId int64) ([]*GetByLivestreamProductId, error) {
	stmt := table.LivestreamExternalVariant.SELECT(
		table.LivestreamExternalVariant.IDLivestreamExternalVariant,
		table.ExternalShop.FkEcommerce,
		table.LivestreamExternalVariant.Quantity,
		table.Variant.IDVariant,
		table.Variant.Option,
		table.ExternalVariant.Price,
	).FROM(
		table.LivestreamExternalVariant.
			INNER_JOIN(table.LivestreamProduct, table.LivestreamProduct.IDLivestreamProduct.EQ(table.LivestreamExternalVariant.FkLivestreamProduct)).
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.LivestreamExternalVariant.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)),
	).WHERE(table.LivestreamExternalVariant.FkLivestreamProduct.EQ(postgres.Int(livestreamProductId)))
	var data []*GetByLivestreamProductId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
