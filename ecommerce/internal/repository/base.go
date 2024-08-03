package repository

import (
	"database/sql"
	"ecommerce/internal/database"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IBaseRepository[T interface{}] interface {
	insertOne(db qrm.Queryable, stmt postgres.Statement) (*T, error)
	insertMany(db qrm.Queryable, stmt postgres.Statement) ([]*T, error)

	CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data T) (*T, error)
	CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*T) ([]*T, error)
	GetById(db qrm.Queryable, id int64) (*T, error)
	ExecWithinTransaction(funcTx func(qrm.Queryable) (interface{}, error)) (interface{}, error)
	GetDefaultDatabase() *database.PostgresqlDatabase
}

type BaseRepository[T interface{}] struct {
	Database *database.PostgresqlDatabase
}

func (r *BaseRepository[T]) insertOne(db qrm.Queryable, stmt postgres.Statement) (*T, error) {
	var data T
	err := stmt.Query(db, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *BaseRepository[T]) insertMany(db qrm.Queryable, stmt postgres.Statement) ([]*T, error) {
	var data []*T
	err := stmt.Query(db, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *BaseRepository[T]) ExecWithinTransaction(funcTx func(db qrm.Queryable) (interface{}, error)) (res interface{}, err error) {
	var tx *sql.Tx
	tx, err = r.Database.Db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			xerr := tx.Rollback() // err is non-nil; don't change it
			if xerr != nil {
				err = errors.Join(err, xerr)
			}
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	res, err = funcTx(tx)
	return
}

func (r *BaseRepository[T]) GetDefaultDatabase() *database.PostgresqlDatabase {
	return r.Database
}
