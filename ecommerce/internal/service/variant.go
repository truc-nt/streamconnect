package service

import (
	"ecommerce/internal/constants"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"

	"github.com/jackc/pgtype"
)

type IVariantService interface {
	GetVariantsByProductId(productId int64, limit int64, offset int64) ([]*GetVariantsByProductId, error)
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

type GetVariantsByProductId struct {
	IDVariant       int64       `json:"id_variant"`
	Sku             string      `json:"sku"`
	Status          string      `json:"status"`
	Option          pgtype.JSON `json:"option"`
	ExternalProduct []*struct {
		IDEcommerce       int16   `json:"id_ecommerce"`
		IDExternalProduct int64   `json:"id_external_product"`
		Ecommerce         string  `json:"ecommerce"`
		Price             float64 `json:"price"`
		Stock             int32   `json:"stock"`
	} `json:"external_products"`
}

func (s *VariantService) GetVariantsByProductId(shopId int64, limit int64, offset int64) ([]*GetVariantsByProductId, error) {
	variants, err := s.VariantRepository.GetVariantsByProductId(s.VariantRepository.GetDefaultDatabase().Db, shopId, limit, offset)
	if err != nil {
		return nil, err
	}
	variantIds := make([]int64, 0)
	for _, variant := range variants {
		variantIds = append(variantIds, variant.IDVariant)
	}
	externalProductShopify, err := s.EcommerceService[constants.SHOPIFY_ID].GetExternalProductByVariantIds(variantIds)
	if err != nil {
		return nil, err
	}

	res := make([]*GetVariantsByProductId, 0)
	for _, variant := range variants {
		for _, externalProduct := range externalProductShopify.([]*model.ExternalProductShopify) {
			if variant.IDVariant == *externalProduct.FkVariant {
				res = append(res, &GetVariantsByProductId{
					IDVariant: variant.IDVariant,
					Sku:       variant.Sku,
					Status:    variant.Status,
					Option:    variant.Option,
					ExternalProduct: []*struct {
						IDEcommerce       int16   `json:"id_ecommerce"`
						IDExternalProduct int64   `json:"id_external_product"`
						Ecommerce         string  `json:"ecommerce"`
						Price             float64 `json:"price"`
						Stock             int32   `json:"stock"`
					}{
						{
							IDEcommerce:       constants.SHOPIFY_ID,
							IDExternalProduct: externalProduct.IDExternalProductShopify,
							Ecommerce:         constants.SHOPIFY_NAME,
							Price:             *externalProduct.Price,
							Stock:             *externalProduct.Stock,
						},
					},
				})
			}
		}
	}

	return res, err
}
