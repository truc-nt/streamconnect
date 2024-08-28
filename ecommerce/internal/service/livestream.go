package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/database/model"
	"ecommerce/internal/database/table"
	"ecommerce/internal/repository"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
)

type ILivestreamService interface {
	CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error
}

type LivestreamService struct {
	LivestreamRepository                repository.ILivestreamRepository
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
}

func NewLivestreamService(livestreamService repository.ILivestreamRepository, livestreamProductRepository repository.ILivestreamProductRepository, livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository) ILivestreamService {
	return &LivestreamService{
		LivestreamRepository:                livestreamService,
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
	}
}

func (s *LivestreamService) CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		newLivestreamData := internalModel.Livestream{
			FkShop:      shopId,
			Title:       createLivestreamRequest.Title,
			Description: &createLivestreamRequest.Description,
			StartTime:   time.Now(),
		}
		newLivestream, err := s.LivestreamRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Livestream.FkShop,
				table.Livestream.Title,
				table.Livestream.Description,
				table.Livestream.StartTime,
			},
			newLivestreamData,
		)
		if err != nil {
			return nil, err
		}

		for _, livestreamProduct := range createLivestreamRequest.LivestreamProducts {
			newLivestreamProductData := internalModel.LivestreamProduct{
				FkLivestream: newLivestream.IDLivestream,
				FkProduct:    livestreamProduct.IDProduct,
				Priority:     livestreamProduct.Priority,
			}
			newLivestreamProduct, err := s.LivestreamProductRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.LivestreamProduct.FkLivestream,
					table.LivestreamProduct.FkProduct,
					table.LivestreamProduct.Priority,
				},
				newLivestreamProductData,
			)
			if err != nil {
				return nil, err
			}

			newExternalLivestreamVariantData := make([]*internalModel.LivestreamExternalVariant, 0)
			for _, livestreamVariant := range livestreamProduct.LivestreamVariants {
				livestreamExternalVariants := lo.Map(livestreamVariant.LivestreamExternalVariants, func(externalVariant *struct {
					IDExternalVariant int64 `json:"id_external_variant"`
					Quantity          int32 `json:"quantity"`
				}, index int) *internalModel.LivestreamExternalVariant {
					return &internalModel.LivestreamExternalVariant{
						FkLivestreamProduct: newLivestreamProduct.IDLivestreamProduct,
						FkExternalVariant:   externalVariant.IDExternalVariant,
						Quantity:            externalVariant.Quantity,
					}
				})
				newExternalLivestreamVariantData = append(newExternalLivestreamVariantData, livestreamExternalVariants...)
			}

			_, err = s.LivestreamExternalVariantRepository.CreateMany(
				db,
				postgres.ColumnList{
					table.LivestreamExternalVariant.FkLivestreamProduct,
					table.LivestreamExternalVariant.FkExternalVariant,
					table.LivestreamExternalVariant.Quantity,
				},
				newExternalLivestreamVariantData,
			)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.LivestreamRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}
