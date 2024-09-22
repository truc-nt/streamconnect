package service

import (
	"ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
	"github.com/samber/lo"
)

type ILivestreamProductService interface {
	GetLivestreamProductsByLivestreamId(livestreamId int64) (interface{}, error)
	GetLivestreamProductInfoByLivestreamProductId(livestreamProductId int64) (interface{}, error)
	PinLivestreamProduct(pinLivestreamProductRequest *model.PinLivestreamProductRequest) error
}

type LivestreamProductService struct {
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
	LivestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository

	EcommerceService map[int16]IEcommerceService
}

func NewLivestreamProductService(
	livestreamProductRepository repository.ILivestreamProductRepository,
	livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository,
	livestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository,
	ecommerceService map[int16]IEcommerceService,
) ILivestreamProductService {
	return &LivestreamProductService{
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
		LivestreamProductFollowerRepository: livestreamProductFollowerRepository,
		EcommerceService:                    ecommerceService,
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

	externalShopsByEcommerceId := make(map[int16]map[int64][]string, 0)
	for _, livestreamVariant := range livestreamExternalVariants.LivestreamVariants {
		for _, livestreamExternalVariant := range livestreamVariant.LivestreamExternalVariants {
			if _, ok := externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce]; !ok {
				externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce] = make(map[int64][]string, 0)
			}

			if _, ok := externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce][livestreamExternalVariant.IDExternalShop]; !ok {
				externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce][livestreamExternalVariant.IDExternalShop] = make([]string, 0)
			}

			if lo.Contains(externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce][livestreamExternalVariant.IDExternalShop], livestreamExternalVariant.ExternalProductIdMapping) {
				continue
			}

			externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce][livestreamExternalVariant.IDExternalShop] = append(externalShopsByEcommerceId[livestreamExternalVariant.IDEcommerce][livestreamExternalVariant.IDExternalShop], livestreamExternalVariant.ExternalProductIdMapping)
		}
	}

	for ecommerceId, productsByExternalShop := range externalShopsByEcommerceId {
		for externalShopId, externalProductIdMappings := range productsByExternalShop {
			externalVariantStocks, err := s.EcommerceService[ecommerceId].GetStockByExternalProductExternalId(externalShopId, externalProductIdMappings)
			if err != nil {
				return nil, err
			}

			for _, externalVariantStock := range externalVariantStocks {
				for _, livestreamVariant := range livestreamExternalVariants.LivestreamVariants {
					for _, livestreamExternalVariant := range livestreamVariant.LivestreamExternalVariants {
						if livestreamExternalVariant.ExternalIdMapping == externalVariantStock.ExternalIdMapping {
							livestreamExternalVariant.Stock += externalVariantStock.Stock
						}
					}
				}
			}
		}
	}

	return livestreamExternalVariants, nil
}

func (s *LivestreamProductService) PinLivestreamProduct(pinLivestreamProductRequest *model.PinLivestreamProductRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, request := range *pinLivestreamProductRequest {
			if _, err := s.LivestreamProductRepository.UpdateById(
				db,
				postgres.ColumnList{
					table.LivestreamProduct.Priority,
				},
				entity.LivestreamProduct{
					IDLivestreamProduct: request.LivestreamProductId,
					Priority:            request.Priority,
				},
			); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.LivestreamProductRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}
