package service

import (
	"ecommerce/api/model"
	"ecommerce/internal/repository"
	"errors"
)

type IProductService interface {
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
