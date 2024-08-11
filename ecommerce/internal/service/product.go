package service

import (
	"ecommerce/api/model"
	internalModel "ecommerce/internal/model"
	"ecommerce/internal/repository"
	"ecommerce/internal/table"
	"encoding/json"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgtype"
)

type IProductService interface {
	GetProductById(productId int64) (*internalModel.Product, error)
	GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error)
	CreateProductVariantsFromExternalProducts(productsVariantsList *model.CreateProductsVariantsRequest) error
}

type ProductService struct {
	ProductRepository         repository.IProductRepository
	VariantRepository         repository.IVariantRepository
	ExternalVariantRepository repository.IExternalVariantRepository

	EcommerceService map[int16]IEcommerceService
}

func NewProductService(productRepository repository.IProductRepository, variantRepository repository.IVariantRepository, externalVariantRepository repository.IExternalVariantRepository, ecommerceService map[int16]IEcommerceService) IProductService {
	return &ProductService{
		ProductRepository:         productRepository,
		VariantRepository:         variantRepository,
		ExternalVariantRepository: externalVariantRepository,
		EcommerceService:          ecommerceService,
	}
}

func (s *ProductService) GetProductById(productId int64) (*internalModel.Product, error) {
	return s.ProductRepository.GetById(s.ProductRepository.GetDefaultDatabase().Db, productId)
}

func (s *ProductService) GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error) {
	return s.ProductRepository.GetByShopId(s.ProductRepository.GetDefaultDatabase().Db, shopId, limit, offset)

}

func (s *ProductService) CreateProductVariantsFromExternalProducts(productsVariantsList *model.CreateProductsVariantsRequest) error {
	/*externalProducts := make(map[int16][]string)
	for _, productsVariant := range *productsVariantsList {
		externalProducts[productsVariant.EcommerceID] = append(externalProducts[productsVariant.EcommerceID], productsVariant.ExternalProductExternalID)
	}

	for ecommerceID, externalProductsList := range externalProducts {
		ecommerceService, ok := s.EcommerceService[ecommerceID]
		if !ok {
			return errors.New("ecommerce service not found")
		}

		if err := ecommerceService.CreateExternalVariants(externalProductsList); err != nil {
			return err
		}
	}

	return nil*/
	idExternalProductList := make([]string, 0)
	for _, productsVariant := range *productsVariantsList {
		idExternalProductList = append(idExternalProductList, productsVariant.IDExternalProduct)
	}

	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		for _, externalProductId := range idExternalProductList {
			externalVariants, err := s.ExternalVariantRepository.GetByExternalProductId(s.ExternalVariantRepository.GetDefaultDatabase().Db, externalProductId)
			if err != nil {
				return nil, err
			}
			if len(externalVariants) == 0 {
				fmt.Printf("external product id %s not found", externalProductId)
				continue
			}

			if externalVariants[0].FkVariant != nil {
				fmt.Printf("shopify priduct id %s already been created", externalProductId)
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
					FkShop: 1,
				},
			)
			if err != nil {
				return nil, err
			}

			var options = make(map[string][]string)

			for _, externalProduct := range externalVariants {
				newVariant, err := s.VariantRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.Variant.FkProduct,
						table.Variant.Sku,
						table.Variant.Option,
					},
					internalModel.Variant{
						FkProduct: newProduct.IDProduct,
						Sku:       externalProduct.Sku,
						Option:    externalProduct.Option,
					},
				)

				if err != nil {
					return nil, err
				}

				var variantOption map[string]string
				if err := json.Unmarshal(externalProduct.Option.Bytes, &variantOption); err != nil {
					return nil, err
				}

				for key, value := range variantOption {
					options[key] = append(options[key], value)
				}

				//updateProductVariantList = append(updateProductVariantList, []int64{newProduct.IDProduct, newVariant.IDVariant, externalProduct.ShopifyVariantID})

				if err := s.ExternalVariantRepository.UpdateExternalVariant(db, newVariant.IDVariant, externalProduct.IDExternal); err != nil {
					return nil, err
				}
			}

			marshalOptions, err := json.Marshal(options)
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
			)

		}

		return nil, nil
	}

	if _, err := s.ProductRepository.ExecWithinTransaction(execWithinTransaction); err != nil {
		return err
	}

	return nil
}
