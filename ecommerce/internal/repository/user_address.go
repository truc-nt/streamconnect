package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IUserAddressRepository interface {
	IBaseRepository[model.UserAddress]
	GetByUserId(db qrm.Queryable, userId int64) ([]*model.UserAddress, error)
	GetDefaultAddressByUserId(db qrm.Queryable, userId int64) (*model.UserAddress, error)
}

type UserAddressRepository struct {
	BaseRepository[model.UserAddress]
}

func NewUserAddressRepository(database *database.PostgresqlDatabase) IUserAddressRepository {
	repo := &UserAddressRepository{}
	repo.Database = database
	return repo
}

func (r *UserAddressRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.UserAddress) (*model.UserAddress, error) {
	stmt := table.UserAddress.INSERT(columnList).MODEL(data).RETURNING(table.UserAddress.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *UserAddressRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.UserAddress) ([]*model.UserAddress, error) {
	stmt := table.UserAddress.INSERT(columnList).MODELS(data).RETURNING(table.UserAddress.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *UserAddressRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.UserAddress) (*model.UserAddress, error) {
	stmt := table.UserAddress.UPDATE(columnList).MODEL(data).RETURNING(table.UserAddress.AllColumns)
	return r.update(db, stmt)
}

func (r *UserAddressRepository) GetById(db qrm.Queryable, id int64) (*model.UserAddress, error) {
	stmt := table.UserAddress.SELECT(table.UserAddress.AllColumns).WHERE(table.UserAddress.IDUserAddress.EQ(postgres.Int(id)))
	var data model.UserAddress
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *UserAddressRepository) GetByUserId(db qrm.Queryable, userId int64) ([]*model.UserAddress, error) {
	stmt := table.UserAddress.SELECT(table.UserAddress.AllColumns).WHERE(table.UserAddress.FkUser.EQ(postgres.Int(userId)))
	var data []*model.UserAddress
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *UserAddressRepository) GetDefaultAddressByUserId(db qrm.Queryable, userId int64) (*model.UserAddress, error) {
	stmt := table.UserAddress.SELECT(table.UserAddress.AllColumns).WHERE(table.UserAddress.FkUser.EQ(postgres.Int(userId)).AND(table.UserAddress.IsDefault.EQ(postgres.Bool(true))))
	var data model.UserAddress
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
