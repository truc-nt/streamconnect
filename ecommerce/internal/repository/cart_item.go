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
		IDLivestreamExternalVariant int64       `alias:"cart_item_livestream_external_variant.id" json:"id_livestream_external_variant"`
		IDEcommerce                 int16       `alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
		Price                       float64     `alias:"external_variant.price" json:"price"`
		Quantity                    int32       `alias:"cart_item.quantity" json:"quantity"`
		MaxQuantity                 int32       `alias:"livestream_external_variant.quantity" json:"max_quantity"`
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

	stmt := table.CartItemLivestreamExternalVariant.SELECT(
		table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,

		table.CartItemLivestreamExternalVariant.AllColumns,

		table.CartItem.IDCartItem,
		table.CartItem.Quantity,
		table.ExternalShop.FkEcommerce,
		table.ExternalVariant.Price,
		imageUrlSubQuery,

		table.LivestreamExternalVariant.Quantity,
	).FROM(
		table.CartItem.
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.CartItem.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)).
			LEFT_JOIN(table.CartItemLivestreamExternalVariant, table.CartItemLivestreamExternalVariant.FkCartItem.EQ(table.CartItem.IDCartItem)).
			LEFT_JOIN(table.LivestreamExternalVariant, table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(table.CartItemLivestreamExternalVariant.Fk)),
	).WHERE(table.CartItem.FkCart.EQ(postgres.Int(cartId)).AND(table.CartItem.Status.EQ(postgres.String(constants.ACTIVE))))

	data := make(GetByCartId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetCartItemsByIds []struct {
	IDEcommerce          int16 `sql:"primary_key" alias:"external_shop.fk_ecommerce" json:"id_ecommerce"`
	CartItemsGroupByShop []*struct {
		IDShop    int64  `sql:"primary_key" alias:"shop.id_shop" json:"id_shop"`
		ShopName  string `alias:"shop.name" json:"shop_name"`
		CartItems []*struct {
			IDCartItem                  int64       `alias:"cart_item.id_cart_item" json:"id_cart_item"`
			Name                        string      `alias:"product.name" json:"name"`
			Option                      pgtype.JSON `alias:"variant.option" json:"option"`
			IDExternalVariant           int64       `alias:"external_variant.id_external_variant" json:"id_external_variant"`
			ExternalIDMapping           string      `alias:"external_variant.external_id_mapping" json:"external_id_mapping"`
			IDLivestreamExternalVariant int64       `alias:"cart_item_livestream_external_variant.id" json:"id_livestream_external_variant"`
			Price                       float64     `alias:"external_variant.price" json:"price"`
			Quantity                    int32       `alias:"cart_item.quantity" json:"quantity"`
			ImageURL                    string      `alias:"image_variant.url" json:"image_url"`
			IDExternalShop              int64       `alias:"external_shop.id_external_shop" json:"id_external_shop"`
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

	stmt := table.CartItemLivestreamExternalVariant.SELECT(
		table.Shop.IDShop,
		table.Shop.Name,

		table.Product.Name,
		table.Variant.Option,

		table.CartItemLivestreamExternalVariant.AllColumns,

		table.CartItem.IDCartItem,
		table.CartItem.Quantity,
		table.ExternalShop.IDExternalShop,
		table.ExternalShop.FkEcommerce,
		table.ExternalVariant.IDExternalVariant,
		table.ExternalVariant.ExternalIDMapping,
		table.ExternalVariant.Price,
		imageUrlSubQuery,

		table.LivestreamExternalVariant.Quantity,
	).FROM(
		table.CartItem.
			INNER_JOIN(table.ExternalVariant, table.ExternalVariant.IDExternalVariant.EQ(table.CartItem.FkExternalVariant)).
			INNER_JOIN(table.ExternalShop, table.ExternalShop.IDExternalShop.EQ(table.ExternalVariant.FkExternalShop)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExternalVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Product.FkShop)).
			LEFT_JOIN(table.CartItemLivestreamExternalVariant, table.CartItemLivestreamExternalVariant.FkCartItem.EQ(table.CartItem.IDCartItem)).
			LEFT_JOIN(table.LivestreamExternalVariant, table.LivestreamExternalVariant.IDLivestreamExternalVariant.EQ(table.CartItemLivestreamExternalVariant.Fk)),
	).WHERE(table.CartItem.IDCartItem.IN(ids...))

	data := make(GetCartItemsByIds, 0)
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

	var data []*model.CartItem
	err := stmt.Query(db, &data)

	if err != nil {
		return err
	}

	return nil
}
