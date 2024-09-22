package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ICartItemLivestreamExternalVariantRepository interface {
	IBaseRepository[model.CartItemLivestreamExtVariant]

	GetByLivestreamExternalVariantIdAndCartId(db qrm.Queryable, livestreamExternalVariantId int64, cartId int64) (*model.CartItemLivestreamExtVariant, error)
}

type CartItemLivestreamExternalVariantRepository struct {
	BaseRepository[model.CartItemLivestreamExtVariant]
}

func NewCartItemLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) ICartItemLivestreamExternalVariantRepository {
	repo := &CartItemLivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *CartItemLivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.CartItemLivestreamExtVariant, error) {
	stmt := table.CartItemLivestreamExtVariant.SELECT(table.CartItemLivestreamExtVariant.AllColumns).WHERE(table.CartItemLivestreamExtVariant.ID.EQ(postgres.Int(int64(id))))

	var data model.CartItemLivestreamExtVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CartItemLivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItemLivestreamExtVariant) (*model.CartItemLivestreamExtVariant, error) {
	stmt := table.CartItemLivestreamExtVariant.INSERT(columnList).MODEL(data).RETURNING(table.CartItemLivestreamExtVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.CartItemLivestreamExtVariant) ([]*model.CartItemLivestreamExtVariant, error) {
	stmt := table.CartItemLivestreamExtVariant.INSERT(columnList).MODELS(data).RETURNING(table.CartItemLivestreamExtVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItemLivestreamExtVariant) (*model.CartItemLivestreamExtVariant, error) {
	stmt := table.CartItemLivestreamExtVariant.UPDATE(columnList).MODEL(data).WHERE(table.CartItemLivestreamExtVariant.ID.EQ(postgres.Int(data.ID))).RETURNING(table.CartItemLivestreamExtVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) GetByLivestreamExternalVariantIdAndCartId(db qrm.Queryable, livestreamExternalVariantId int64, cartId int64) (*model.CartItemLivestreamExtVariant, error) {
	stmt := table.CartItemLivestreamExtVariant.SELECT(table.CartItemLivestreamExtVariant.AllColumns).
		FROM(
			table.CartItemLivestreamExtVariant.
				INNER_JOIN(table.CartItem, table.CartItemLivestreamExtVariant.FkCartItem.EQ(table.CartItem.IDCartItem)),
		).WHERE(table.CartItemLivestreamExtVariant.FkLivestreamExtVariant.EQ(postgres.Int(int64(livestreamExternalVariantId))).AND(table.CartItem.FkCart.EQ(postgres.Int(int64(cartId)))))

	var data model.CartItemLivestreamExtVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
