package repository

import (
	apiModel "ecommerce/api/model"
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
)

type ILivestreamRepository interface {
	IBaseRepository[model.Livestream]
	GetByParam(db qrm.Queryable, param *apiModel.GetLivestreamsQueryParam) ([]*model.Livestream, error)

	GetInfoById(db qrm.Queryable, id int64) (*GetInfo, error)
	GetOrdersByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*model.OrderItem, error)
}

type LivestreamRepository struct {
	BaseRepository[model.Livestream]
}

func NewLivestreamRepository(database *database.PostgresqlDatabase) ILivestreamRepository {
	repo := &LivestreamRepository{}
	repo.Database = database
	return repo
}

func (r *LivestreamRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Livestream) (*model.Livestream, error) {
	stmt := table.Livestream.INSERT(columnList).MODEL(data).RETURNING(table.Livestream.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *LivestreamRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Livestream) ([]*model.Livestream, error) {
	stmt := table.Livestream.INSERT(columnList).MODELS(data).RETURNING(table.Livestream.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *LivestreamRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Livestream) (*model.Livestream, error) {
	stmt := table.Livestream.UPDATE(columnList).MODEL(data).WHERE(table.Livestream.IDLivestream.EQ(postgres.Int(data.IDLivestream))).RETURNING(table.Livestream.AllColumns)
	return r.update(db, stmt)
}

func (r *LivestreamRepository) GetById(db qrm.Queryable, id int64) (*model.Livestream, error) {
	stmt := table.Livestream.SELECT(table.Livestream.AllColumns).WHERE(table.Livestream.IDLivestream.EQ(postgres.Int(int64(id))))

	var data model.Livestream
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *LivestreamRepository) GetByParam(db qrm.Queryable, param *apiModel.GetLivestreamsQueryParam) ([]*model.Livestream, error) {
	conditions := postgres.Bool(true)

	if len(param.Status) > 0 {
		statuses := lo.Map(param.Status, func(status string, _ int) postgres.Expression {
			return postgres.String(status)
		})

		conditions = conditions.AND(table.Livestream.Status.IN(statuses...))
	}
	if param.ShopId != 0 {
		conditions = conditions.AND(table.Livestream.FkShop.EQ(postgres.Int(int64(param.ShopId))))
	}

	stmt := table.Livestream.SELECT(table.Livestream.AllColumns).WHERE(conditions)

	data := make([]*model.Livestream, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetInfo struct {
	*model.Livestream
	ShopName string `alias:"shop.name" json:"shop_name"`
}

func (r *LivestreamRepository) GetInfoById(db qrm.Queryable, id int64) (*GetInfo, error) {
	stmt := table.Livestream.SELECT(
		table.Livestream.AllColumns,
		table.Shop.Name,
	).FROM(
		table.Livestream.
			LEFT_JOIN(table.Shop, table.Livestream.FkShop.EQ(table.Shop.IDShop)),
	).WHERE(table.Livestream.IDLivestream.EQ(postgres.Int(int64(id))))

	var data GetInfo
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *LivestreamRepository) GetOrdersByLivestreamId(db qrm.Queryable, livestreamId int64) ([]*model.OrderItem, error) {
	stmt := table.ExtOrder.SELECT(
		table.OrderItem.AllColumns,
	).FROM(
		table.OrderItemLivestreamExtVariant.
			INNER_JOIN(table.LivestreamExtVariant, table.LivestreamExtVariant.IDLivestreamExtVariant.EQ(table.OrderItemLivestreamExtVariant.FkLivestreamExtVariant)).
			INNER_JOIN(table.LivestreamProduct, table.LivestreamProduct.IDLivestreamProduct.EQ(table.LivestreamExtVariant.FkLivestreamProduct)).
			INNER_JOIN(table.OrderItem, table.OrderItem.IDOrderItem.EQ(table.OrderItemLivestreamExtVariant.FkOrderItem)).
			INNER_JOIN(table.Order, table.Order.IDOrder.EQ(table.OrderItem.FkOrder)).
			INNER_JOIN(table.ExtOrder, table.ExtOrder.FkOrder.EQ(table.ExtOrder.FkOrder)),
	).WHERE(table.LivestreamProduct.FkLivestream.EQ(postgres.Int(livestreamId)))

	data := make([]*model.OrderItem, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
