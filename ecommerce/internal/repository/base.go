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
	update(db qrm.Queryable, stmt postgres.Statement) (*T, error)

	CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data T) (*T, error)
	CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*T) ([]*T, error)
	UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data T) (*T, error)
	GetById(db qrm.Queryable, id int64) (*T, error)
	ExecWithinTransaction(funcTx func(qrm.Queryable) (interface{}, error)) (interface{}, error)
	GetDatabase() *database.PostgresqlDatabase
}

type BaseRepository[T interface{}] struct {
	Database *database.PostgresqlDatabase
}

/*func (r *BaseRepository[T]) Exec(fn func(db qrm.Queryable) (interface{}, error)) (interface{}, error) {
	return fn(r.Database.Db)
}

func (r *BaseRepository[T]) ExecTx(db qrm.Queryable, fn func(db qrm.Queryable) (interface{}, error)) (interface{}, error) {
	return fn(db)
}*/

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

func (r *BaseRepository[T]) update(db qrm.Queryable, stmt postgres.Statement) (*T, error) {
	var data T
	err := stmt.Query(db, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *BaseRepository[T]) ExecWithinTransaction(fnTx func(db qrm.Queryable) (interface{}, error)) (interface{}, error) {
	var tx *sql.Tx
	var err error
	tx, err = r.Database.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			xerr := tx.Rollback()
			if xerr != nil {
				err = errors.Join(err, xerr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	res, err := fnTx(tx)
	return res, err
}

func (r *BaseRepository[T]) GetDatabase() *database.PostgresqlDatabase {
	return r.Database
}
