package service

import (
	"ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IExternalVariantService interface {
	GetExternalVariantsGroupByProduct(limit int64, offset int64) (interface{}, error)
	GetExternalVariantsByExternalProductIdMapping(externalProductIdMapping string) ([]*entity.ExtVariant, error)
	ConnectVariants(connectVariantsRequest *model.ConnectVariantsRequest) error
}

type ExternalVariantService struct {
	VariantRepository         repository.IVariantRepository
	ExternalVariantRepository repository.IExternalVariantRepository
}

func NewExternalVariantService(externalVariantRepository repository.IExternalVariantRepository, variantRepository repository.IVariantRepository) IExternalVariantService {
	return &ExternalVariantService{
		ExternalVariantRepository: externalVariantRepository,
		VariantRepository:         variantRepository,
	}
}

func (s *ExternalVariantService) GetExternalVariantsGroupByProduct(limit int64, offset int64) (interface{}, error) {
	externalProducts, err := s.ExternalVariantRepository.GetExternalVariantsGroupByProduct(s.ExternalVariantRepository.GetDatabase().Db, limit, offset)
	if err != nil {
		return nil, err
	}

	return externalProducts, nil
}

func (s *ExternalVariantService) GetExternalVariantsByExternalProductIdMapping(externalProductIdMapping string) ([]*entity.ExtVariant, error) {
	externalProducts, err := s.ExternalVariantRepository.GetExternalVariantsByExternalProductIdMapping(s.ExternalVariantRepository.GetDatabase().Db, externalProductIdMapping)
	if err != nil {
		return nil, err
	}

	return externalProducts, nil
}

func (s *ExternalVariantService) ConnectVariants(connectVariantsRequest *model.ConnectVariantsRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, connectVariants := range *connectVariantsRequest {
			variant, err := s.VariantRepository.GetVariantInfoById(s.VariantRepository.GetDatabase().Db, connectVariants.IDVariant)
			if err != nil {
				return nil, err
			}

			externalVariant, err := s.ExternalVariantRepository.GetExternalVariantInfoById(s.ExternalVariantRepository.GetDatabase().Db, connectVariants.IDExternalVariant)
			if err != nil {
				return nil, err
			}

			if variant.IDShop != externalVariant.IDShop {
				return nil, errors.New("variant and external variant are not from the same shop")
			}

			if _, err := s.ExternalVariantRepository.UpdateById(
				db,
				postgres.ColumnList{
					table.ExtVariant.FkVariant,
				},
				entity.ExtVariant{
					IDExtVariant: connectVariants.IDExternalVariant,
					FkVariant:    &connectVariants.IDVariant,
				},
			); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.ExternalVariantRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}
	return nil

}
