package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	"github.com/samber/lo"
)

type ICartService interface {
	AddToCart(cartId int64, CartItemList []*model.AddToCartRequest) error
	Get(cartId int64) (interface{}, error)
	Update(cartItemId int64, updateRequest *model.UpdateRequest) error

	GetCartItemsByIdsRequest(getCartItemsByIdsRequest []*model.GetCartItemsByIdsRequest) (interface{}, error)
}

type CartService struct {
	CartItemRepository                repository.ICartItemRepository
	LivestreamExternalVariant         repository.ILivestreamExternalVariantRepository
	CartItemLivestreamExternalVariant repository.ICartItemLivestreamExternalVariantRepository
}

func NewCartService(
	cartItemRepository repository.ICartItemRepository,
	livestreamExternalVariant repository.ILivestreamExternalVariantRepository,
	cartItemLivestreamExternalVariant repository.ICartItemLivestreamExternalVariantRepository,
) ICartService {
	return &CartService{
		CartItemRepository:                cartItemRepository,
		LivestreamExternalVariant:         livestreamExternalVariant,
		CartItemLivestreamExternalVariant: cartItemLivestreamExternalVariant,
	}
}

func (s *CartService) Get(cartId int64) (interface{}, error) {
	return s.CartItemRepository.GetByCartId(s.CartItemRepository.GetDatabase().Db, cartId)
}

func (s *CartService) AddToCart(cartId int64, addToCartRequest []*model.AddToCartRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, cartItem := range addToCartRequest {
			findCartItem, err := s.CartItemLivestreamExternalVariant.GetByLivestreamExternalVariantIdAndCartId(db, cartItem.IDLivestreamExternalVariant, cartId)
			var pgErr pgx.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code != pgerrcode.NoDataFound {
					return nil, err
				}
			}

			if findCartItem != nil {
				_, err := s.CartItemRepository.UpdateById(
					db,
					postgres.ColumnList{
						table.CartItem.Quantity,
					},
					internalModel.CartItem{
						IDCartItem: findCartItem.FkCartItem,
						Quantity:   cartItem.Quantity,
					},
				)
				if err != nil {
					return nil, err
				}
				continue
			}

			livestreamExternalVariant, err := s.LivestreamExternalVariant.GetById(db, cartItem.IDLivestreamExternalVariant)
			if err != nil {
				return nil, err
			}

			newCartItem, err := s.CartItemRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.CartItem.FkCart,
					table.CartItem.FkExternalVariant,
					table.CartItem.Quantity,
				},
				internalModel.CartItem{
					FkCart:            cartId,
					FkExternalVariant: livestreamExternalVariant.FkExternalVariant,
					Quantity:          cartItem.Quantity,
				},
			)
			if err != nil {
				return nil, err
			}

			_, err = s.CartItemLivestreamExternalVariant.CreateOne(
				db,
				postgres.ColumnList{
					table.CartItemLivestreamExternalVariant.FkCartItem,
					table.CartItemLivestreamExternalVariant.Fk,
				},
				internalModel.CartItemLivestreamExternalVariant{
					FkCartItem: newCartItem.IDCartItem,
					Fk:         cartItem.IDLivestreamExternalVariant,
				},
			)

			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.CartItemRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *CartService) Update(cartItemId int64, updateRequest *model.UpdateRequest) error {
	_, err := s.CartItemRepository.UpdateById(
		s.CartItemRepository.GetDatabase().Db,
		postgres.ColumnList{
			table.CartItem.Quantity,
		},
		internalModel.CartItem{
			IDCartItem: cartItemId,
			Quantity:   updateRequest.Quantity,
		})
	return err
}

func (s *CartService) GetCartItemsByIdsRequest(getCartItemsByIdsRequest []*model.GetCartItemsByIdsRequest) (interface{}, error) {
	cartItemIds := lo.Map(getCartItemsByIdsRequest, func(request *model.GetCartItemsByIdsRequest, _ int) int64 {
		return request.IDCartItem
	})

	return s.CartItemRepository.GetCartItemsByIds(s.CartItemRepository.GetDatabase().Db, cartItemIds)
}
