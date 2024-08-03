package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/model"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IProductRepository interface {
	IBaseRepository[model.Product]
}

type ProductRepository struct {
	BaseRepository[model.Product]
}

func NewProductRepository(database *database.PostgresqlDatabase) IProductRepository {
	repo := &ProductRepository{}
	repo.Database = database
	return repo
}

func (r *ProductRepository) GetById(db qrm.Queryable, id int64) (*model.Product, error) {
	stmt := table.Product.SELECT(table.Product.AllColumns).WHERE(table.Product.IDProduct.EQ(postgres.Int(int64(id))))

	var data model.Product
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *ProductRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.Product) (*model.Product, error) {
	stmt := table.Product.INSERT(columnList).MODEL(data).RETURNING(table.Product.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *ProductRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.Product) ([]*model.Product, error) {
	stmt := table.Product.INSERT(columnList).MODELS(data).RETURNING(table.Product.AllColumns)
	return r.insertMany(db, stmt)
}
