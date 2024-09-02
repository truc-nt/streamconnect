package service

import (
	"ecommerce/internal/repository"

	"github.com/samber/lo"
)

type IVariantService interface {
	GetVariantsByProductId(productId int64, limit int64, offset int64) ([]*repository.GetVariantsByProductId, error)

	GetVariantsByExternalProductIdMapping(externalProductIdMapping string) (interface{}, error)
}

type VariantService struct {
	VariantRepository repository.IVariantRepository

	EcommerceService map[int16]IEcommerceService
}

func NewVariantService(variantRepository repository.IVariantRepository, ecommerceService map[int16]IEcommerceService) IVariantService {
	return &VariantService{
		VariantRepository: variantRepository,
		EcommerceService:  ecommerceService,
	}
}

func (s *VariantService) GetVariantsByProductId(shopId int64, limit int64, offset int64) ([]*repository.GetVariantsByProductId, error) {
	variants, err := s.VariantRepository.GetVariantsByProductId(s.VariantRepository.GetDatabase().Db, shopId, limit, offset)
	if err != nil {
		return nil, err
	}

	externalShopsByEcommerceId := make(map[int16]map[int64][]string, 0)
	for _, variant := range variants {
		for _, externalVariant := range variant.ExternalVariants {
			if _, ok := externalShopsByEcommerceId[externalVariant.IDEcommerce]; !ok {
				externalShopsByEcommerceId[externalVariant.IDEcommerce] = make(map[int64][]string, 0)
			}

			if _, ok := externalShopsByEcommerceId[externalVariant.IDEcommerce][externalVariant.IDExternalShop]; !ok {
				externalShopsByEcommerceId[externalVariant.IDEcommerce][externalVariant.IDExternalShop] = make([]string, 0)
			}

			if lo.Contains(externalShopsByEcommerceId[externalVariant.IDEcommerce][externalVariant.IDExternalShop], externalVariant.ExternalProductIdMapping) {
				continue
			}

			externalShopsByEcommerceId[externalVariant.IDEcommerce][externalVariant.IDExternalShop] = append(externalShopsByEcommerceId[externalVariant.IDEcommerce][externalVariant.IDExternalShop], externalVariant.ExternalProductIdMapping)
		}
	}

	for ecommerceId, productsByExternalShop := range externalShopsByEcommerceId {
		for externalShopId, externalProductIdMappings := range productsByExternalShop {
			externalVariantStocks, err := s.EcommerceService[ecommerceId].GetStockByExternalProductExternalId(externalShopId, externalProductIdMappings)
			if err != nil {
				return nil, err
			}

			for _, externalVariantStock := range externalVariantStocks {
				for _, variant := range variants {
					for _, externalVariant := range variant.ExternalVariants {
						if externalVariant.ExternalIdMapping == externalVariantStock.ExternalIdMapping {
							externalVariant.Stock += externalVariantStock.Stock
						}
					}
				}
			}
		}
	}

	//externalVariantStocks, err := s.EcommerceService[constants.SHOPIFY_ID].GetStockByExternalProductExternalId(shopId, externalProductExternalId)

	return variants, err
}

func (s *VariantService) GetVariantsByExternalProductIdMapping(externalProductIdMapping string) (interface{}, error) {
	return s.VariantRepository.GetVariantsByExternalProductIdMapping(s.VariantRepository.GetDatabase().Db, externalProductIdMapping)
}
