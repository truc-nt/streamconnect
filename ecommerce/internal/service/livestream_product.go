package service

import (
	"ecommerce/internal/repository"

	"github.com/jackc/pgtype"
	"github.com/samber/lo"
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
	livestreamProducts, err := s.LivestreamProductRepository.GetByLivestreamId(s.LivestreamProductRepository.GetDefaultDatabase().Db, livestreamId)
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

type GetLivestreamProductInfoByLivestreamProductId struct {
	Name               string                                        `json:"name"`
	Description        string                                        `json:"description"`
	Option             pgtype.JSON                                   `json:"option"`
	LivestreamVariants []*GetLivestreamVariantsByLivestreamProductId `json:"livestream_variants"`
}

func (s *LivestreamProductService) GetLivestreamProductInfoByLivestreamProductId(livestreamProductId int64) (interface{}, error) {
	livesetreamProduct, err := s.LivestreamProductRepository.GetInfoById(s.LivestreamProductRepository.GetDefaultDatabase().Db, livestreamProductId)
	if err != nil {
		return nil, err
	}

	livestreamExternalVariants, err := s.LivestreamExternalVariantRepository.GetByLivestreamProductId(s.LivestreamExternalVariantRepository.GetDefaultDatabase().Db, livestreamProductId)
	if err != nil {
		return nil, err
	}
	livestreamVariantsMapping := make(map[int64]*GetLivestreamVariantsByLivestreamProductId)
	for _, livestreamExternalVariant := range livestreamExternalVariants {
		if _, ok := livestreamVariantsMapping[livestreamExternalVariant.IDLivestreamExternalVariant]; !ok {
			livestreamVariantsMapping[livestreamExternalVariant.IDLivestreamExternalVariant] = &GetLivestreamVariantsByLivestreamProductId{
				Option: livestreamExternalVariant.Option,
			}
		}

		livestreamVariantsMapping[livestreamExternalVariant.IDLivestreamExternalVariant].LivestreamExternalVariants = append(livestreamVariantsMapping[livestreamExternalVariant.IDLivestreamExternalVariant].LivestreamExternalVariants, &GetLivestreamExternalVariantByLivestreamProductId{
			IDLivestreamExternalVariant: livestreamExternalVariant.IDLivestreamExternalVariant,
			IDEcommerce:                 livestreamExternalVariant.IDEcommerce,
			Quantity:                    livestreamExternalVariant.Quantity,
			Price:                       livestreamExternalVariant.Price,
		})
	}
	if err != nil {
		return nil, err
	}

	livestreamVariants := lo.Values(livestreamVariantsMapping)
	res := &GetLivestreamProductInfoByLivestreamProductId{
		Name:               livesetreamProduct.Name,
		Description:        livesetreamProduct.Description,
		Option:             livesetreamProduct.Option,
		LivestreamVariants: livestreamVariants,
	}

	return res, nil
}
