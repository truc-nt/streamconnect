package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IExternalOrderRepository interface {
	IBaseRepository[model.ExtOrder]
	GetExternalOrdersByOrderId(db qrm.Queryable, orderId int64) (*GetExternalOrdersByOrderId, error)
}

type ExternalOrderRepository struct {
	BaseRepository[model.ExtOrder]
}

func NewExternalOrderRepository(database *database.PostgresqlDatabase) IExternalOrderRepository {
	repo := &ExternalOrderRepository{}
	repo.Database = database
	return repo
}

func (r *ExternalOrderRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtOrder) (*model.ExtOrder, error) {
	stmt := table.ExtOrder.INSERT(columnList).MODEL(data).RETURNING(table.ExtOrder.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ExternalOrderRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ExtOrder) ([]*model.ExtOrder, error) {
	stmt := table.ExtOrder.INSERT(columnList).MODELS(data).RETURNING(table.ExtOrder.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *ExternalOrderRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ExtOrder) (*model.ExtOrder, error) {
	stmt := table.ExtOrder.UPDATE(columnList).MODEL(data).WHERE(table.ExtOrder.IDExtOrder.EQ(postgres.Int(data.IDExtOrder))).RETURNING(table.ExtOrder.AllColumns)
	return r.update(db, stmt)
}

func (r *ExternalOrderRepository) GetById(db qrm.Queryable, id int64) (*model.ExtOrder, error) {
	stmt := table.ExtOrder.SELECT(table.ExtOrder.AllColumns).WHERE(table.ExtOrder.IDExtOrder.EQ(postgres.Int(int64(id))))

	var data model.ExtOrder
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type GetExternalOrdersByOrderId []*struct {
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
}

func (r *ExternalOrderRepository) GetExternalOrdersByOrderId(db qrm.Queryable, orderId int64) (*GetExternalOrdersByOrderId, error) {
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
	).WHERE(table.Order.IDOrder.EQ(postgres.Int(orderId)))

	data := make(GetExternalOrdersByOrderId, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
