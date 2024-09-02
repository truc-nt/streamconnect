package service

import (
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/repository"
)

type IExternalShopAuthService interface {
	GetShopifyAuthByExternalShopId(externalShopId int64) (*model.ExternalShopShopifyAuth, error)
}

type ExternalShopAuthService struct {
	ExternalShopShopifyAuthRepository repository.IExternalShopShopifyAuthRepository
}

func NewExternalShopAuthService(repo repository.IExternalShopShopifyAuthRepository) IExternalShopAuthService {
	return &ExternalShopAuthService{
		ExternalShopShopifyAuthRepository: repo,
	}
}

func (s *ExternalShopAuthService) GetShopifyAuthByExternalShopId(externalShopId int64) (*model.ExternalShopShopifyAuth, error) {
	auth, err := s.ExternalShopShopifyAuthRepository.GetByExternalShopId(s.ExternalShopShopifyAuthRepository.GetDatabase().Db, externalShopId)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
