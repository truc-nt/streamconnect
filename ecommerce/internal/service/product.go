package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/database/model"
	"ecommerce/internal/database/table"
	"ecommerce/internal/repository"
	"encoding/json"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IProductService interface {
	GetProductById(productId int64) (*internalModel.Product, error)
	GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error)
	CreateProductWithVariants(shopId int64, createProductsWithVariantsRequest *model.CreateProductWithVariants) error
}

type ProductService struct {
	ProductRepository         repository.IProductRepository
	VariantRepository         repository.IVariantRepository
	ExternalVariantRepository repository.IExternalVariantRepository
	ImageVariantRepository    repository.IImageVariantRepository

	EcommerceService map[int16]IEcommerceService
}

func NewProductService(
	productRepository repository.IProductRepository,
	variantRepository repository.IVariantRepository,
	externalVariantRepository repository.IExternalVariantRepository,
	imageVariantRepository repository.IImageVariantRepository,
	ecommerceService map[int16]IEcommerceService) IProductService {
	return &ProductService{
		ProductRepository:         productRepository,
		VariantRepository:         variantRepository,
		ExternalVariantRepository: externalVariantRepository,
		ImageVariantRepository:    imageVariantRepository,
		EcommerceService:          ecommerceService,
	}
}

func (s *ProductService) GetProductById(productId int64) (*internalModel.Product, error) {
	return s.ProductRepository.GetById(s.ProductRepository.GetDatabase().Db, productId)
}

func (s *ProductService) GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error) {
	return s.ProductRepository.GetByShopId(s.ProductRepository.GetDatabase().Db, shopId, limit, offset)

}

func (s *ProductService) CreateProductWithVariants(shopId int64, createProductsWithVariantsRequest *model.CreateProductWithVariants) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		for _, createProductsWithVariants := range *createProductsWithVariantsRequest {
			externalVariants, err := s.ExternalVariantRepository.GetByExternalProductId(s.ExternalVariantRepository.GetDatabase().Db, createProductsWithVariants.ExternalProductIdMapping)
			if err != nil {
				return nil, err
			}
			if len(externalVariants) == 0 {
				fmt.Printf("external product id %s not found", createProductsWithVariants.ExternalProductIdMapping)
				continue
			}

			if externalVariants[0].FkVariant != nil {
				fmt.Printf("shopify priduct id %s already been created", createProductsWithVariants.ExternalProductIdMapping)
				continue
			}

			newProduct, err := s.ProductRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.Product.Name,
					table.Product.FkShop,
				},
				internalModel.Product{
					Name:   externalVariants[0].Name,
					FkShop: shopId,
				},
			)
			if err != nil {
				return nil, err
			}

			for _, externalVariant := range externalVariants {
				newVariant, err := s.VariantRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.Variant.FkProduct,
						table.Variant.Sku,
						table.Variant.Option,
						table.Variant.Status,
					},
					internalModel.Variant{
						FkProduct: newProduct.IDProduct,
						Sku:       externalVariant.Sku,
						Option:    externalVariant.Option,
						Status:    externalVariant.Status,
					},
				)

				if err != nil {
					return nil, err
				}

				var variantOption map[string]string
				if err := json.Unmarshal(externalVariant.Option.Bytes, &variantOption); err != nil {
					return nil, err
				}

				//updateProductVariantList = append(updateProductVariantList, []int64{newProduct.IDProduct, newVariant.IDVariant, externalProduct.ShopifyVariantID})

				if err := s.ExternalVariantRepository.UpdateExternalVariant(db, newVariant.IDVariant, externalVariant.ExternalIDMapping); err != nil {
					return nil, err
				}

				if _, err := s.ImageVariantRepository.CreateOne(db, postgres.ColumnList{
					table.ImageVariant.FkVariant,
					table.ImageVariant.URL,
				}, internalModel.ImageVariant{
					FkVariant: newVariant.IDVariant,
					URL:       *externalVariant.ImageURL,
				}); err != nil {
					return nil, err
				}
			}

			/*marshalOptions, err := json.Marshal(options)
			if err != nil {
				return nil, err
			}

			s.ProductRepository.UpdateById(
				db,
				postgres.ColumnList{table.Product.Option},
				internalModel.Product{
					IDProduct: newProduct.IDProduct,
					Option: pgtype.JSON{
						Bytes:  marshalOptions,
						Status: pgtype.Present,
					},
				},
			)*/

		}

		return nil, nil
	}

	if _, err := s.ProductRepository.ExecWithinTransaction(execWithinTransaction); err != nil {
		return err
	}

	return nil
}
