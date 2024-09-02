package service

import (
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/repository"
)

type IUserService interface {
	GetDefaultAddressByUserId(userId int64) (*model.UserAddress, error)
}

type UserService struct {
	UserAddressRepository repository.IUserAddressRepository
}

func NewUserService(userAddressRepository repository.IUserAddressRepository) IUserService {
	return &UserService{
		UserAddressRepository: userAddressRepository,
	}
}

func (s *UserService) GetDefaultAddressByUserId(userId int64) (*model.UserAddress, error) {
	return s.UserAddressRepository.GetDefaultAddressByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
}
