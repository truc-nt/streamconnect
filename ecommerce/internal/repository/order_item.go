package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IOrderItemRepository interface {
	IBaseRepository[model.OrderItem]
	GetByUserId(db qrm.Queryable, userId int64) (*GetByUserId, error)
}

type OrderItemRepository struct {
	BaseRepository[model.OrderItem]
}

func NewOrderItemRepository(database *database.PostgresqlDatabase) IOrderItemRepository {
	repo := &OrderItemRepository{}
	repo.Database = database
	return repo
}

func (r *OrderItemRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItem) (*model.OrderItem, error) {
	stmt := table.OrderItem.INSERT(columnList).MODEL(data).RETURNING(table.OrderItem.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *OrderItemRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.OrderItem) ([]*model.OrderItem, error) {
	stmt := table.OrderItem.INSERT(columnList).MODELS(data).RETURNING(table.OrderItem.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *OrderItemRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.OrderItem) (*model.OrderItem, error) {
	stmt := table.OrderItem.UPDATE(columnList).MODEL(data).WHERE(table.OrderItem.IDOrderItem.EQ(postgres.Int(data.IDOrderItem))).RETURNING(table.OrderItem.AllColumns)
	return r.update(db, stmt)
}

func (r *OrderItemRepository) GetById(db qrm.Queryable, id int64) (*model.OrderItem, error) {
	stmt := table.OrderItem.SELECT(table.OrderItem.AllColumns).WHERE(table.OrderItem.IDOrderItem.EQ(postgres.Int(int64(id))))

	var data model.OrderItem
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

/*type GetByUserId []*struct {
	IDOrder        int64 `sql:"primary_key" alias:"order.id_order" json:"id_order"`
	ExternalOrders []*struct {
		IDExternalOrder        int64   `sql:"primary_key" alias:"ext_order.id_external_order" json:"id_external_order"`
		ExternalOrderIdMapping string  `alias:"ext_order.ext_order_id_mapping" json:"external_order_id_mapping"`
		Subtotal               float64 `alias:"ext_order.subtotal" json:"subtotal"`
		ShippingFee            float64 `alias:"ext_order.shipping_fee" json:"shipping_fee"`
		InternalDiscount       float64 `alias:"ext_order.internal_discount" json:"internal_discount"`
		ExternalDiscount       float64 `alias:"ext_order.external_discount" json:"external_discount"`
		OrderItems             []*struct {
			IDOrderItem       int64 `sql:"primary_key" alias:"order_item.id_order_item" json:"id_order_item"`
			FkExternalVariant int64 `alias:"order_item.fk_ext_variant" json:"fk_external_variant"`
			Quantity          int64 `alias:"order_item.quantity" json:"quantity"`
			UnitPrice         int64 `alias:"order_item.unit_price" json:"unit_price"`
			PaidPrice         int64 `alias:"order_item.paid_price" json:"paid_price"`
		} `json:"order_items"`
	} `json:"external_orders"`
}*/

type GetByUserId []*struct {
	IDOrder        int64  `sql:"primary_key" alias:"order.id_order" json:"id_order"`
	IDShop         int64  `alias:"order.fk_shop" json:"id_shop"`
	ShopName       string `alias:"shop.name" json:"shop_name"`
	ExternalOrders []*struct {
		IDExternalOrder        int64   `sql:"primary_key" alias:"ext_order.id_ext_order" json:"id_external_order"`
		IDEcommerce            int64   `alias:"ext_shop.fk_ecommerce" json:"id_ecommerce"`
		ExternalOrderIdMapping string  `alias:"ext_order.ext_order_id_mapping" json:"external_order_id_mapping"`
		ShippingFee            float64 `alias:"ext_order.shipping_fee" json:"shipping_fee"`
		InternalDiscount       float64 `alias:"ext_order.internal_discount" json:"internal_discount"`
		ExternalDiscount       float64 `alias:"ext_order.external_discount" json:"external_discount"`
		OrderItems             []*struct {
			IDOrderItem       int64       `sql:"primary_key" alias:"order_item.id_order_item" json:"id_order_item"`
			Name              string      `alias:"product.name" json:"name"`
			FkExternalVariant int64       `alias:"order_item.fk_ext_variant" json:"fk_external_variant"`
			Quantity          int64       `alias:"order_item.quantity" json:"quantity"`
			UnitPrice         float64     `alias:"order_item.unit_price" json:"unit_price"`
			PaidPrice         float64     `alias:"order_item.paid_price" json:"paid_price"`
			Option            pgtype.JSON `alias:"variant.option" json:"option"`
			ImageURL          string      `alias:"image_variant.url" json:"image_url"`
		} `json:"order_items"`
	} `json:"external_orders"`
}

func (r *OrderItemRepository) GetByUserId(db qrm.Queryable, userId int64) (*GetByUserId, error) {
	innerVariantAlias := table.Variant.AS("inner_variant")
	imageUrlSubQuery := table.ImageVariant.SELECT(table.ImageVariant.URL).
		FROM(
			table.ImageVariant.
				LEFT_JOIN(innerVariantAlias, innerVariantAlias.IDVariant.EQ(table.ImageVariant.FkVariant)),
		).WHERE(table.Variant.IDVariant.EQ(innerVariantAlias.IDVariant)).LIMIT(1)

	stmt := table.OrderItem.SELECT(
		table.Order.AllColumns,
		table.Shop.Name,
		table.ExtOrder.AllColumns,
		table.OrderItem.AllColumns,
		table.Variant.AllColumns,
		table.ExtShop.AllColumns,
		table.Product.Name,
		imageUrlSubQuery,
	).FROM(
		table.OrderItem.
			INNER_JOIN(table.Order, table.Order.IDOrder.EQ(table.OrderItem.FkOrder)).
			INNER_JOIN(table.Shop, table.Shop.IDShop.EQ(table.Order.FkShop)).
			INNER_JOIN(table.ExtVariant, table.ExtVariant.IDExtVariant.EQ(table.OrderItem.FkExtVariant)).
			INNER_JOIN(table.Variant, table.Variant.IDVariant.EQ(table.ExtVariant.FkVariant)).
			INNER_JOIN(table.Product, table.Product.IDProduct.EQ(table.Variant.FkProduct)).
			INNER_JOIN(table.ExtShop, table.ExtShop.IDExtShop.EQ(table.ExtVariant.FkExtShop)).
			INNER_JOIN(table.ExtOrder, table.ExtOrder.FkExtShop.EQ(table.ExtShop.IDExtShop).AND(table.ExtOrder.FkOrder.EQ(table.Order.IDOrder))),
	).WHERE(table.Order.FkUser.EQ(postgres.Int(userId)))
	data := make(GetByUserId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
