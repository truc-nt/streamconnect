package service

import (
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/repository"
)

type IUserService interface {
	GetByUserId(userId int64) (*model.User, error)
	GetDefaultAddressByUserId(userId int64) (*model.UserAddress, error)
}

type UserService struct {
	UserRepository        repository.IUserRepository
	UserAddressRepository repository.IUserAddressRepository
}

func NewUserService(
	userRepository repository.IUserRepository,
	userAddressRepository repository.IUserAddressRepository,
) IUserService {
	return &UserService{
		UserRepository:        userRepository,
		UserAddressRepository: userAddressRepository,
	}
}

func (s *UserService) GetDefaultAddressByUserId(userId int64) (*model.UserAddress, error) {
	return s.UserAddressRepository.GetDefaultAddressByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
}

func (s *UserService) GetByUserId(userId int64) (*model.User, error) {
	return s.UserRepository.GetById(s.UserAddressRepository.GetDatabase().Db, userId)
}
