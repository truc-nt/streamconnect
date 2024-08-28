package service

import (
	"ecommerce/internal/database/model"
	"ecommerce/internal/repository"
)

type IExternalVariantService interface {
	GetExternalVariantsGroupByProduct(limit int64, offset int64) (interface{}, error)
	GetExternalVariantsByExternalProductIdMapping(externalProductIdMapping string) ([]*model.ExternalVariant, error)
}

type ExternalVariantService struct {
	Repository repository.IExternalVariantRepository
}

func NewExternalVariantService(repo repository.IExternalVariantRepository) IExternalVariantService {
	return &ExternalVariantService{
		Repository: repo,
	}
}

func (s *ExternalVariantService) GetExternalVariantsGroupByProduct(limit int64, offset int64) (interface{}, error) {
	externalProducts, err := s.Repository.GetExternalVariantsGroupByProduct(s.Repository.GetDatabase().Db, limit, offset)
	if err != nil {
		return nil, err
	}

	return externalProducts, nil
}

func (s *ExternalVariantService) GetExternalVariantsByExternalProductIdMapping(externalProductIdMapping string) ([]*model.ExternalVariant, error) {
	externalProducts, err := s.Repository.GetExternalVariantsByExternalProductIdMapping(s.Repository.GetDatabase().Db, externalProductIdMapping)
	if err != nil {
		return nil, err
	}

	return externalProducts, nil
}
