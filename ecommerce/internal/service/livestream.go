package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/table"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
)

type ILivestreamService interface {
	CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error
	GetProductsByLivestreamId(livestreamId int64) (interface{}, error)
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

		livestreamProductsMapping := make(map[int64]bool, 0)
		for _, product := range createLivestreamRequest.LivestreamExternalVariants {
			livestreamProductsMapping[product.IDProduct] = true
		}

		livestreamProducts := lo.Map(lo.Keys(livestreamProductsMapping), func(productId int64, index int) *internalModel.LivestreamProduct {
			var priority int32 = int32(index)
			return &internalModel.LivestreamProduct{
				FkLivestream: newLivestream.IDLivestream,
				FkProduct:    productId,
				Priority:     &priority,
			}
		})

		newLivestreamProductList, err := s.LivestreamProductRepository.CreateMany(
			db,
			postgres.ColumnList{
				table.LivestreamProduct.FkLivestream,
				table.LivestreamProduct.FkProduct,
				table.LivestreamProduct.Priority,
			},
			livestreamProducts,
		)
		if err != nil {
			return nil, err
		}

		fmt.Println(newLivestreamProductList)

		livestreamExternalVariants := lo.Map(createLivestreamRequest.LivestreamExternalVariants, func(externalVariant *model.CreateLivestreamExternalVariants, index int) *internalModel.LivestreamExternalVariant {
			livestreamProduct, ok := lo.Find(newLivestreamProductList, func(livestreamProduct *internalModel.LivestreamProduct) bool {
				return livestreamProduct.FkProduct == externalVariant.IDProduct
			})

			if !ok {
				return nil
			}

			return &internalModel.LivestreamExternalVariant{
				FkLivestreamProduct: livestreamProduct.IDLivestreamProduct,
				FkExternalVariant:   externalVariant.IDExternalVariant,
				Quantity:            externalVariant.Quantity,
			}
		})

		_, err = s.LivestreamExternalVariantRepository.CreateMany(
			db,
			postgres.ColumnList{
				table.LivestreamExternalVariant.FkLivestreamProduct,
				table.LivestreamExternalVariant.FkExternalVariant,
				table.LivestreamExternalVariant.Quantity,
			},
			livestreamExternalVariants,
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err := s.LivestreamRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *LivestreamService) GetProductsByLivestreamId(livestreamId int64) (interface{}, error) {
	livestreamProducts, err := s.LivestreamProductRepository.GetByLivestreamId(s.LivestreamProductRepository.GetDefaultDatabase().Db, livestreamId)
	if err != nil {
		return nil, err
	}

	return livestreamProducts, nil
}
