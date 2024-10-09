package service

import (
	apiModel "ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"

	"github.com/go-jet/jet/v2/postgres"
)

type IUserAddressService interface {
	GetAddressByUserId(userId int64) ([]*entity.UserAddress, error)
	CreateAddress(userId int64, request *apiModel.CreateAddressRequest) (*entity.UserAddress, error)
}

type UserAddressService struct {
	UserAddressRepository repository.IUserAddressRepository
}

func NewUserAddressService(
	UserAddressRepository repository.IUserAddressRepository,
) IUserAddressService {
	return &UserAddressService{
		UserAddressRepository: UserAddressRepository,
	}
}

func (s *UserAddressService) GetAddressByUserId(userId int64) ([]*entity.UserAddress, error) {
	return s.UserAddressRepository.GetByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
}

func (s *UserAddressService) CreateAddress(userId int64, request *apiModel.CreateAddressRequest) (*entity.UserAddress, error) {
	addresses, err := s.UserAddressRepository.GetByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
	if err != nil {
		return nil, err
	}

	isDefault := false
	if len(addresses) == 0 {
		isDefault = true
	}

	address := &entity.UserAddress{
		FkUser:    userId,
		Name:      request.Name,
		Phone:     request.Phone,
		Address:   request.Address,
		City:      request.City,
		IsDefault: isDefault,
	}

	return s.UserAddressRepository.CreateOne(
		s.UserAddressRepository.GetDatabase().Db,
		postgres.ColumnList{
			table.UserAddress.FkUser,
			table.UserAddress.Name,
			table.UserAddress.Phone,
			table.UserAddress.Address,
			table.UserAddress.City,
			table.UserAddress.IsDefault,
		},
		*address,
	)
}
