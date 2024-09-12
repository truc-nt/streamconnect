package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IVoucherUserRepository interface {
	IBaseRepository[model.VoucherUser]
}

type VoucherUserRepository struct {
	BaseRepository[model.VoucherUser]
}

func NewVoucherUserRepository(database *database.PostgresqlDatabase) IVoucherUserRepository {
	repo := &VoucherUserRepository{}
	repo.Database = database
	return repo
}

func (r *VoucherUserRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.VoucherUser) (*model.VoucherUser, error) {
	stmt := table.VoucherUser.
		INSERT(columnList).
		MODEL(data).
		RETURNING(table.VoucherUser.AllColumns).
		ON_CONFLICT(
			table.VoucherUser.FkVoucher,
			table.VoucherUser.FkUser,
		).
		DO_NOTHING()
	return r.insertOne(db, stmt)
}

func (r *VoucherUserRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.VoucherUser) ([]*model.VoucherUser, error) {
	stmt := table.VoucherUser.INSERT(columnList).MODELS(data).RETURNING(table.VoucherUser.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *VoucherUserRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.VoucherUser) (*model.VoucherUser, error) {
	stmt := table.VoucherUser.UPDATE(columnList).MODEL(data).RETURNING(table.VoucherUser.AllColumns)
	return r.update(db, stmt)
}

func (r *VoucherUserRepository) GetById(db qrm.Queryable, id int64) (*model.VoucherUser, error) {
	stmt := table.VoucherUser.SELECT(table.VoucherUser.AllColumns).WHERE(table.VoucherUser.IDVoucherUser.EQ(postgres.Int(id)))
	var data model.VoucherUser
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
