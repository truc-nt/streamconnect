package repository

import (
	"ecommerce/internal/constants"
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
	"github.com/samber/lo"
)

type ICartItemRepository interface {
	IBaseRepository[model.CartItem]
	GetByCartId(db qrm.Queryable, cartId int64) (*GetByCartId, error)
	GetCartItemsByIds(db qrm.Queryable, cartItemIds []int64) (*GetCartItemsByIds, error)
	UpdateCartItemsToInactiveByIds(db qrm.Queryable, cartItemIds []int64) error

	GetCartItemInfoById(db qrm.Queryable, id int64) (*GetCartItemInfoById, error)
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
	IDShop   int64  `sql:"primary_key" alias:"shop.id_shop" json:"id_shop"`
	ShopName string `alias:"shop.name" json:"shop_name"`
	CartItem []*struct {
		IDCartItem                  int64       `alias:"cart_item.id_cart_item" json:"id_cart_item"`
		Name                        string      `alias:"product.name" json:"name"`
		Option                      pgtype.JSON `alias:"variant.option" json:"option"`
		IDLivestreamExternalVariant int64       `alias:"cart_item_livestream_ext_variant.id" json:"id_livestream_external_variant"`
		IDEcommerce                 int16       `alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
		Price                       float64     `alias:"ext_variant.price" json:"price"`
		Quantity                    int32       `alias:"cart_item.quantity" json:"quantity"`
		MaxQuantity                 int32       `alias:"livestream_ext_variant.quantity" json:"max_quantity"`
		ImageURL                    string      `alias:"image_variant.url" json:"image_url"`
	} `json:"cart_items"`
}

func (r *CartItemRepository) GetByCartId(db qrm.Queryable, cartId int64) (*GetByCartId, error) {
	innerVariantAlias := table.Variant.AS("inner_variant")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(innerVariantAlias, innerVariantAlias.IDVariant.EQ(table.ImageVariant.FkVariant)),
		).WHERE(table.Variant.IDVariant.EQ(innerVariantAlias.IDVariant)).LIMIT(1)

	stmt := table.CartItemLivestreamExtVariant.SELECT(
		table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,

		table.CartItemLivestreamExtVariant.AllColumns,

		table.CartItem.IDCartItem,
		table.CartItem.Quantity,
		table.ExtShop.FkEcommerce,
		table.ExtVariant.Price,
		imageUrlSubQuery,

		table.LivestreamExtVariant.Quantity,
	).FROM(
		table.CartItem.
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.CartItem.FkExtVariant)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)).
			LEFT_JOIN(table.CartItemLivestreamExtVariant, table.CartItemLivestreamExtVariant.FkCartItem.EQ(table.CartItem.IDCartItem)).
			LEFT_JOIN(table.LivestreamExtVariant, table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(table.CartItemLivestreamExtVariant.FkLivestreamExtVariant)),
	).WHERE(table.CartItem.FkCart.EQ(postgres.Int(cartId)).AND(table.CartItem.Status.EQ(postgres.String(constants.ACTIVE))))

	data := make(GetByCartId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetCartItemsByIds []struct {
	IDEcommerce          int16 `sql:"primary_key" alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
	CartItemsGroupByShop []*struct {
		IDShop    int64  `sql:"primary_key" alias:"shop.id_shop" json:"id_shop"`
		ShopName  string `alias:"shop.name" json:"shop_name"`
		CartItems []*struct {
			IDCartItem                  int64       `alias:"cart_item.id_cart_item" json:"id_cart_item"`
			Name                        string      `alias:"product.name" json:"name"`
			Option                      pgtype.JSON `alias:"variant.option" json:"option"`
			IDExternalVariant           int64       `alias:"ext_variant.id_ext_variant" json:"id_external_variant"`
			ExternalIDMapping           string      `alias:"ext_variant.ext_id_mapping" json:"external_id_mapping"`
			IDLivestreamExternalVariant int64       `alias:"cart_item_livestream_ext_variant.fk_livestream_ext_variant" json:"id_livestream_external_variant"`
			Price                       float64     `alias:"ext_variant.price" json:"price"`
			Quantity                    int32       `alias:"cart_item.quantity" json:"quantity"`
			ImageURL                    string      `alias:"image_variant.url" json:"image_url"`
			IDExternalShop              int64       `alias:"ext_shop.id_ext_shop" json:"id_external_shop"`
		} `json:"cart_items"`
	} `json:"cart_items_group_by_shop"`
}

func (r *CartItemRepository) GetCartItemsByIds(db qrm.Queryable, cartItemIds []int64) (*GetCartItemsByIds, error) {
	ids := lo.Map(cartItemIds, func(id int64, _ int) postgres.Expression {
		return postgres.Int(id)
	})

	innerVariantAlias := table.Variant.AS("inner_variant")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(innerVariantAlias, innerVariantAlias.IDVariant.EQ(table.ImageVariant.FkVariant)),
		).WHERE(table.Variant.IDVariant.EQ(innerVariantAlias.IDVariant)).LIMIT(1)

	stmt := table.CartItemLivestreamExtVariant.SELECT(
		table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,

		table.CartItemLivestreamExtVariant.AllColumns,

		table.CartItem.IDCartItem,
		table.CartItem.Quantity,
		table.ExtShop.IDExtShop,
		table.ExtShop.FkEcommerce,
		table.ExtVariant.IDExtVariant,
		table.ExtVariant.ExtIDMapping,
		table.ExtVariant.Price,
		imageUrlSubQuery,

		table.LivestreamExtVariant.Quantity,
	).FROM(
		table.CartItem.
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.CartItem.FkExtVariant)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)).
			LEFT_JOIN(table.CartItemLivestreamExtVariant, table.CartItemLivestreamExtVariant.FkCartItem.EQ(table.CartItem.IDCartItem)).
			LEFT_JOIN(table.LivestreamExtVariant, table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(table.CartItemLivestreamExtVariant.FkLivestreamExtVariant)),
	).WHERE(table.CartItem.IDCartItem.IN(ids...))

	data := make(GetCartItemsByIds, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetCartItemInfoById struct {
	model.CartItem
	IDLivestreamExternalVariant int64   `alias:"cart_item_livestream_ext_variant.fk_livestream_ext_variant" json:"id_livestream_external_variant"`
	IDExternalVariant           int64   `alias:"ext_variant.id_ext_variant" json:"id_external_variant"`
	IDEcommerce                 int16   `alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
	IDExternalShop              int64   `alias:"ext_shop.id_ext_shop" json:"id_external_shop"`
	ExternalIDMapping           string  `alias:"ext_variant.ext_id_mapping" json:"external_id_mapping"`
	Price                       float64 `alias:"ext_variant.price" json:"price"`
}

func (r *CartItemRepository) GetCartItemInfoById(db qrm.Queryable, id int64) (*GetCartItemInfoById, error) {
	stmt := table.CartItem.SELECT(
		table.CartItem.AllColumns,
		table.CartItemLivestreamExtVariant.FkLivestreamExtVariant,
		table.ExtVariant.IDExtVariant,
		table.ExtVariant.ExtIDMapping,
		table.ExtShop.IDExtShop,
		table.ExtShop.FkEcommerce,
		table.ExtVariant.Price,
	).FROM(
		table.CartItem.
			LEFT_JOIN(table.CartItemLivestreamExtVariant, table.CartItemLivestreamExtVariant.FkCartItem.EQ(table.CartItem.IDCartItem)).
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.CartItem.FkExtVariant)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)),
	).WHERE(table.CartItem.IDCartItem.EQ(postgres.Int(int64(id))))

	var data GetCartItemInfoById
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CartItemRepository) UpdateCartItemsToInactiveByIds(db qrm.Queryable, cartItemIds []int64) error {
	ids := lo.Map(cartItemIds, func(id int64, _ int) postgres.Expression {
		return postgres.Int(id)
	})

	stmt := table.CartItem.UPDATE(table.CartItem.Status).SET(postgres.String(constants.INACTIVE)).WHERE(table.CartItem.IDCartItem.IN(ids...)).RETURNING(table.CartItem.AllColumns)

	data := make([]*model.CartItem, 0)
	err := stmt.Query(db, &data)

	if err != nil {
		return err
	}

	return nil
}
