package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/table"

	"ecommerce/internal/database/model"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ICartItemRepository interface {
	IBaseRepository[model.CartItem]
	GetByCartId(db qrm.Queryable, cartId int64) ([]*model.CartItemLivestreamExternalVariant, error)
}

type CartItemRepository struct {
	BaseRepository[model.CartItem]
}

func NewCartItemRepository(database *database.PostgresqlDatabase) ICartItemRepository {
	repo := &CartItemRepository{}
	repo.Database = database
	return repo
}

func (r *CartItemRepository) GetById(db qrm.Queryable, id int64) (*model.CartItem, error) {
	stmt := table.CartItem.SELECT(table.Cart.AllColumns).WHERE(table.CartItem.IDCartItem.EQ(postgres.Int(int64(id))))

	var data model.CartItem
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CartItemRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItem) (*model.CartItem, error) {
	stmt := table.CartItem.INSERT(columnList).MODEL(data).RETURNING(table.CartItem.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *CartItemRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.CartItem) ([]*model.CartItem, error) {
	stmt := table.CartItem.INSERT(columnList).MODELS(data).RETURNING(table.CartItem.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *CartItemRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItem) (*model.CartItem, error) {
	stmt := table.CartItem.UPDATE(columnList).MODEL(data).WHERE(table.CartItem.IDCartItem.EQ(postgres.Int(data.IDCartItem))).RETURNING(table.CartItem.AllColumns)
	return r.update(db, stmt)
}

type GetByCartId []struct {
	/*IDShop   int64  `sql:"primary_key" alias:"shop.id_shop" json:"id_shop"`
	ShopName string `alias:"shop.name" json:"shop_name"`
	CartItem []struct {
		IDCartItem                  int64       `sql:"primary_key" alias:"cart_item.id_cart_item" json:"id_cart_item"`
		Name                        string      `alias:"product.name" json:"name"`
		Option                      pgtype.JSON `alias:"variant.option" json:"option"`
		IDLivestreamExternalVariant int64       `alias:"cart_item_livestream_external_variant.id_cart_item_livestream_external_variant" json:"id_livestream_external_variant"`
		IDEcommerce                 int16       `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
		Price                       float64     `alias:"external_variant.price" json:"price"`
		Quantity                    int32       `alias:"cart_item.quantity" json:"quantity"`
	} `json:"cart_item"`*/

	IDCartItem int64  `sql:"primary_key" alias:"cart_item.id_cart_item" json:"id_cart_item"`
	Name       string `alias:"product.name" json:"name"`
	//Option                      pgtype.JSON `alias:"variant.option" json:"option"`
	IDLivestreamExternalVariant int64   `alias:"cart_item_livestream_external_variant.id_cart_item_livestream_external_variant" json:"id_livestream_external_variant"`
	IDEcommerce                 int16   `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
	Price                       float64 `alias:"external_variant.price" json:"price"`
	Quantity                    int32   `alias:"cart_item.quantity" json:"quantity"`
}

func (r *CartItemRepository) GetByCartId(db qrm.Queryable, cartId int64) ([]*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.SELECT(
		/*table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,*/

		table.CartItemLivestreamExternalVariant.AllColumns,

		//table.CartItem.IDCartItem,
		//table.CartItem.Quantity,
		/*table.ExternalShop.FkEcommerce,
		table.ExternalVariant.Price,*/
	)
	/*.FROM(
		table.CartItem.
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.CartItem.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)).
			INNER_JOIN(table.CartItemLivestreamExternalVariant, table.CartItemLivestreamExternalVariant.FkCartItem.EQ(table.CartItem.IDCartItem)),
	).WHERE(table.CartItem.FkCart.EQ(postgres.Int(cartId)))*/

	var data []*model.CartItemLivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
