package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/table"

	"github.com/go-jet/jet/v2/postgres"
)

type ICartService interface {
	AddToCart(cartId int64, cartLivestreamExternalVariantList *model.AddToCartRequest) error
	Get(cartId int64) (interface{}, error)
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func NewCartService(
	cartRepository repository.ICartRepository,
) ICartService {
	return &CartService{
		CartRepository: cartRepository,
	}
}

func (s *CartService) Get(cartId int64) (interface{}, error) {
	return s.CartRepository.GetByCartId(s.CartRepository.GetDefaultDatabase().Db, cartId)
}

func (s *CartService) AddToCart(cartId int64, cartLivestreamExternalVariantList *model.AddToCartRequest) error {
	cartLivestreamExternalVariants := make([]*internalModel.CartLivestreamExternalVariant, 0, len(*cartLivestreamExternalVariantList))
	for _, cartLivestreamExternalVariant := range *cartLivestreamExternalVariantList {
		cartLivestreamExternalVariants = append(cartLivestreamExternalVariants, &internalModel.CartLivestreamExternalVariant{
			FkCart:                      cartId,
			FkLivestreamExternalVariant: cartLivestreamExternalVariant.IDLivestreamExternalVariant,
			Quantity:                    cartLivestreamExternalVariant.Quantity,
		})
	}

	_, err := s.CartRepository.CreateMany(
		s.CartRepository.GetDefaultDatabase().Db,
		postgres.ColumnList{
			table.CartLivestreamExternalVariant.FkCart,
			table.CartLivestreamExternalVariant.FkLivestreamExternalVariant,
			table.CartLivestreamExternalVariant.Quantity,
		},
		cartLivestreamExternalVariants,
	)
	if err != nil {
		return err
	}

	return nil
}
