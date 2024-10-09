package service

import (
	apiModel "ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"

	"github.com/go-jet/jet/v2/postgres"
)

type IUserService interface {
	GetByUserId(userId int64) (*entity.User, error)
	GetDefaultAddressByUserId(userId int64) (*entity.UserAddress, error)
	UpdateUser(userId int64, user *apiModel.UpdateUserRequest) error
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

func (s *UserService) GetDefaultAddressByUserId(userId int64) (*entity.UserAddress, error) {
	return s.UserAddressRepository.GetDefaultAddressByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
}

func (s *UserService) GetByUserId(userId int64) (*entity.User, error) {
	return s.UserRepository.GetById(s.UserAddressRepository.GetDatabase().Db, userId)
}

func (s *UserService) UpdateUser(userId int64, request *apiModel.UpdateUserRequest) error {
	updatedColumnList := postgres.ColumnList{}
	user := entity.User{
		IDUser: userId,
	}

	if request.Gender != nil {
		updatedColumnList = append(
			updatedColumnList,
			table.User.Gender,
		)
		user.Gender = request.Gender
	}
	if request.Email != nil {
		updatedColumnList = append(
			updatedColumnList,
			table.User.Email,
		)
		user.Email = *request.Email
	}
	if request.Birthdate != nil {
		updatedColumnList = append(
			updatedColumnList,
			table.User.Birthdate,
		)
		user.Birthdate = request.Birthdate
	}

	if _, err := s.UserRepository.UpdateById(
		s.UserRepository.GetDatabase().Db,
		updatedColumnList,
		user,
	); err != nil {
		return err
	}

	return nil
}
