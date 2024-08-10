package service

import (
	"ecommerce/api/model"
	internalModal "ecommerce/internal/model"
	"ecommerce/internal/repository"
	"errors"
)

type IProductService interface {
	GetProductById(productId int64) (*internalModal.Product, error)
	GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error)
	CreateProductsVariantsFromExternalProducts(productsVariantsList *model.CreateProductsVariantsRequest) error
}

type ProductService struct {
	ProductRepository repository.IProductRepository

	EcommerceService map[int16]IEcommerceService
}

func NewProductService(productRepository repository.IProductRepository, ecommerceService map[int16]IEcommerceService) IProductService {
	return &ProductService{
		ProductRepository: productRepository,
		EcommerceService:  ecommerceService,
	}
}

func (s *ProductService) GetProductById(productId int64) (*internalModal.Product, error) {
	return s.ProductRepository.GetById(s.ProductRepository.GetDefaultDatabase().Db, productId)
}

func (s *ProductService) GetProductsByShopId(shopId int64, limit int64, offset int64) (interface{}, error) {
	return s.ProductRepository.GetByShopId(s.ProductRepository.GetDefaultDatabase().Db, shopId, limit, offset)

}

func (s *ProductService) CreateProductsVariantsFromExternalProducts(productsVariantsList *model.CreateProductsVariantsRequest) error {
	externalProducts := make(map[int16][]string)
	for _, productsVariant := range *productsVariantsList {
		externalProducts[productsVariant.EcommerceID] = append(externalProducts[productsVariant.EcommerceID], productsVariant.ExternalProductExternalID)
	}

	for ecommerceID, externalProductsList := range externalProducts {
		ecommerceService, ok := s.EcommerceService[ecommerceID]
		if !ok {
			return errors.New("ecommerce service not found")
		}

		if err := ecommerceService.CreateProductVariants(externalProductsList); err != nil {
			return err
		}
	}

	return nil
}
