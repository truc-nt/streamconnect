package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/table"

	"ecommerce/internal/model"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type ICartRepository interface {
	IBaseRepository[model.CartLivestreamExternalVariant]
	GetByCartId(db qrm.Queryable, cartId int64) (*GetByCartId, error)
}

type CartRepository struct {
	BaseRepository[model.CartLivestreamExternalVariant]
}

func NewCartRepository(database *database.PostgresqlDatabase) ICartRepository {
	repo := &CartRepository{}
	repo.Database = database
	return repo
}

func (r *CartRepository) GetById(db qrm.Queryable, id int64) (*model.CartLivestreamExternalVariant, error) {
	stmt := table.CartLivestreamExternalVariant.SELECT(table.Cart.AllColumns).WHERE(table.CartLivestreamExternalVariant.IDCartLivestreamExternalVariant.EQ(postgres.Int(int64(id))))

	var data model.CartLivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CartRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.CartLivestreamExternalVariant) (*model.CartLivestreamExternalVariant, error) {
	stmt := table.CartLivestreamExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.CartLivestreamExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *CartRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.CartLivestreamExternalVariant) ([]*model.CartLivestreamExternalVariant, error) {
	stmt := table.CartLivestreamExternalVariant.INSERT(columnList).MODELS(data).RETURNING(table.CartLivestreamExternalVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *CartRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.CartLivestreamExternalVariant) (*model.CartLivestreamExternalVariant, error) {
	stmt := table.CartLivestreamExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.CartLivestreamExternalVariant.IDCartLivestreamExternalVariant.EQ(postgres.Int(data.IDCartLivestreamExternalVariant))).RETURNING(table.CartLivestreamExternalVariant.AllColumns)
	return r.update(db, stmt)
}

type GetByCartId []struct {
	IDShop                        int64  `sql:"primary_key" alias:"shop.id_shop" json:"id_shop"`
	ShopName                      string `alias:"shop.name" json:"shop_name"`
	CartLivestreamExternalVariant []struct {
		Name                        string      `alias:"product.name" json:"name"`
		Option                      pgtype.JSON `alias:"variant.option" json:"option"`
		IDLivestreamExternalVariant int64       `alias:"livestream_external_variant.id_livestream_external_variant" json:"id_livestream_external_variant"`
		IDEcommerce                 int16       `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
		Price                       float64     `alias:"external_variant.price" json:"price"`
		Quantity                    int32       `alias:"cart_livestream_external_variant.quantity" json:"quantity"`
	} `json:"cart_livestream_external_variant"`
}

func (r *CartRepository) GetByCartId(db qrm.Queryable, cartId int64) (*GetByCartId, error) {
	stmt := table.CartLivestreamExternalVariant.SELECT(
		table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,

		table.LivestreamExternalVariant.IDLivestreamExternalVariant,
		table.ExternalVariant.Price,
		table.CartLivestreamExternalVariant.Quantity,
		table.ExternalShop.FkEcommerce,
	).FROM(
		table.CartLivestreamExternalVariant.
			INNER_JOIN(table.LivestreamExternalVariant, table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(table.CartLivestreamExternalVariant.FkLivestreamExternalVariant)).
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.LivestreamExternalVariant.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)),
	).WHERE(table.CartLivestreamExternalVariant.FkCart.EQ(postgres.Int(cartId)))

	var data GetByCartId
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
