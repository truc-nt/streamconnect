package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type ICartItemLivestreamExternalVariantRepository interface {
	IBaseRepository[model.CartItemLivestreamExternalVariant]

	GetByLivestreamExternalVariantIdAndCartId(db qrm.Queryable, livestreamExternalVariantId int64, cartId int64) (*model.CartItemLivestreamExternalVariant, error)
}

type CartItemLivestreamExternalVariantRepository struct {
	BaseRepository[model.CartItemLivestreamExternalVariant]
}

func NewCartItemLivestreamExternalVariantRepository(database *database.PostgresqlDatabase) ICartItemLivestreamExternalVariantRepository {
	repo := &CartItemLivestreamExternalVariantRepository{}
	repo.Database = database
	return repo
}

func (r *CartItemLivestreamExternalVariantRepository) GetById(db qrm.Queryable, id int64) (*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.SELECT(table.CartItemLivestreamExternalVariant.AllColumns).WHERE(table.CartItemLivestreamExternalVariant.ID.EQ(postgres.Int(int64(id))))

	var data model.CartItemLivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *CartItemLivestreamExternalVariantRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItemLivestreamExternalVariant) (*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.INSERT(columnList).MODEL(data).RETURNING(table.CartItemLivestreamExternalVariant.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.CartItemLivestreamExternalVariant) ([]*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.INSERT(columnList).MODELS(data).RETURNING(table.CartItemLivestreamExternalVariant.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.CartItemLivestreamExternalVariant) (*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.UPDATE(columnList).MODEL(data).WHERE(table.CartItemLivestreamExternalVariant.ID.EQ(postgres.Int(data.ID))).RETURNING(table.CartItemLivestreamExternalVariant.AllColumns)
	return r.update(db, stmt)
}

func (r *CartItemLivestreamExternalVariantRepository) GetByLivestreamExternalVariantIdAndCartId(db qrm.Queryable, livestreamExternalVariantId int64, cartId int64) (*model.CartItemLivestreamExternalVariant, error) {
	stmt := table.CartItemLivestreamExternalVariant.SELECT(table.CartItemLivestreamExternalVariant.AllColumns).
		FROM(
			table.CartItemLivestreamExternalVariant.
				INNER_JOIN(table.CartItem, table.CartItemLivestreamExternalVariant.FkCartItem.EQ(table.CartItem.IDCartItem)),
		).WHERE(table.CartItemLivestreamExternalVariant.Fk.EQ(postgres.Int(int64(livestreamExternalVariantId))).AND(table.CartItem.FkCart.EQ(postgres.Int(int64(cartId)))))

	var data model.CartItemLivestreamExternalVariant
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
