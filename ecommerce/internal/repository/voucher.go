package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
)

type IVoucherRepository interface {
	IBaseRepository[model.Voucher]
	GetShopVouchers(db qrm.Queryable, userId, shopId int64) ([]*GetShopVouchers, error)
	GetByIds(db qrm.Queryable, ids []int64) ([]*model.Voucher, error)
}

type VoucherRepository struct {
	BaseRepository[model.Voucher]
}

func NewVoucherRepository(database *database.PostgresqlDatabase) IVoucherRepository {
	repo := &VoucherRepository{}
	repo.Database = database
	return repo
}

func (r *VoucherRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Voucher) (*model.Voucher, error) {
	stmt := table.Voucher.INSERT(columnList).MODEL(data).RETURNING(table.Voucher.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *VoucherRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Voucher) ([]*model.Voucher, error) {
	stmt := table.Voucher.INSERT(columnList).MODELS(data).RETURNING(table.Voucher.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *VoucherRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.Voucher) (*model.Voucher, error) {
	stmt := table.Voucher.UPDATE(columnList).MODEL(data).RETURNING(table.Voucher.AllColumns)
	return r.update(db, stmt)
}

func (r *VoucherRepository) GetById(db qrm.Queryable, id int64) (*model.Voucher, error) {
	stmt := table.Voucher.SELECT(table.Voucher.AllColumns).WHERE(table.Voucher.IDVoucher.EQ(postgres.Int(id)))
	var data model.Voucher
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *VoucherRepository) GetByIds(db qrm.Queryable, voucherIds []int64) ([]*model.Voucher, error) {
	ids := lo.Map(voucherIds, func(id int64, _ int) postgres.Expression {
		return postgres.Int(id)
	})

	stmt := table.Voucher.SELECT(table.Voucher.AllColumns).WHERE(table.Voucher.IDVoucher.IN(ids...))
	data := make([]*model.Voucher, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type GetShopVouchers struct {
	*model.Voucher
	IsSaved bool `json:"is_saved"`
}

func (r *VoucherRepository) GetShopVouchers(db qrm.Queryable, userId, shopId int64) ([]*GetShopVouchers, error) {
	stmt := table.VoucherUser.SELECT(
		table.Voucher.AllColumns,
		postgres.CASE().
			WHEN(table.VoucherUser.FkUser.EQ(postgres.Int(userId))).THEN(postgres.Bool(true)).
			ELSE(postgres.Bool(false)).AS("GetShopVouchers.is_saved"),
	).
		FROM(
			table.Voucher.
				LEFT_JOIN(table.VoucherUser, table.Voucher.IDVoucher.EQ(table.VoucherUser.FkVoucher).AND(table.VoucherUser.FkUser.EQ(postgres.Int(userId)))),
		).
		WHERE(table.Voucher.FkShop.EQ(postgres.Int(shopId)))
	data := make([]*GetShopVouchers, 0)
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
