package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IUserRepository interface {
	IBaseRepository[model.User]
}

type UserRepository struct {
	BaseRepository[model.User]
}

func NewUserRepository(database *database.PostgresqlDatabase) IUserRepository {
	repo := &UserRepository{}
	repo.Database = database
	return repo
}

func (r *UserRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.User) (*model.User, error) {
	stmt := table.User.INSERT(columnList).MODEL(data).RETURNING(table.User.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *UserRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.User) ([]*model.User, error) {
	stmt := table.User.INSERT(columnList).MODELS(data).RETURNING(table.User.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *UserRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.User) (*model.User, error) {
	stmt := table.User.UPDATE(columnList).MODEL(data).WHERE(table.User.IDUser.EQ(postgres.Int(data.IDUser))).RETURNING(table.User.AllColumns)
	return r.update(db, stmt)
}

func (r *UserRepository) GetById(db qrm.Queryable, id int64) (*model.User, error) {
	stmt := table.User.SELECT(table.User.AllColumns).WHERE(table.User.IDUser.EQ(postgres.Int(int64(id))))

	var data model.User
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
