package service

import (
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/repository"
)

type IUserAddressService interface {
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
