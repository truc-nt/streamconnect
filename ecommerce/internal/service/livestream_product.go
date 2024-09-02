package service

import (
	"ecommerce/internal/repository"

	"github.com/jackc/pgtype"
)

type ILivestreamProductService interface {
	GetLivestreamProductsByLivestreamId(livestreamId int64) (interface{}, error)
	GetLivestreamProductInfoByLivestreamProductId(livestreamProductId int64) (interface{}, error)
}

type LivestreamProductService struct {
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
}

func NewLivestreamProductService(livestreamProductRepository repository.ILivestreamProductRepository, livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository) ILivestreamProductService {
	return &LivestreamProductService{
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
	}
}

func (s *LivestreamProductService) GetLivestreamProductsByLivestreamId(livestreamId int64) (interface{}, error) {
	livestreamProducts, err := s.LivestreamProductRepository.GetByLivestreamId(s.LivestreamProductRepository.GetDatabase().Db, livestreamId)
	if err != nil {
		return nil, err
	}

	return livestreamProducts, nil
}

type GetLivestreamExternalVariantByLivestreamProductId struct {
	IDLivestreamExternalVariant int64   `json:"id_livestream_external_variant"`
	IDEcommerce                 int16   `json:"id_ecommerce"`
	Quantity                    int64   `json:"quantity"`
	Price                       float64 `json:"price"`
}

type GetLivestreamVariantsByLivestreamProductId struct {
	Option                     pgtype.JSON                                          `json:"option"`
	LivestreamExternalVariants []*GetLivestreamExternalVariantByLivestreamProductId `json:"livestream_external_variants"`
}

func (s *LivestreamProductService) GetLivestreamProductInfoByLivestreamProductId(livestreamProductId int64) (interface{}, error) {
	livestreamExternalVariants, err := s.LivestreamExternalVariantRepository.GetByLivestreamProductId(s.LivestreamExternalVariantRepository.GetDatabase().Db, livestreamProductId)
	if err != nil {
		return nil, err
	}

	return livestreamExternalVariants, nil
}
