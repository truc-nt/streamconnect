package service

import (
	"ecommerce/api/model"
	"ecommerce/internal/constants"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"strconv"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IShopService interface {
	CreateShopForNewAccount(request *model.CreateShopForNewUserRequest) error
	GetShop(shopId int64) (*entity.Shop, error)
	UpdateShop(shopId int64, request *model.UpdateShopRequest) error

	IsFollowed(shopId int64, userId int64) (bool, error)
	AddFollower(shopId int64, userId int64) error
}

type ShopService struct {
	ShopRepository        repository.IShopRepository
	CartRepository        repository.ICartRepository
	AclRoleUserRepository repository.IAclRoleUserRepository
}

func NewShopService(
	shopRepository repository.IShopRepository,
	cartRepository repository.ICartRepository,
	aclRoleUserRepository repository.IAclRoleUserRepository,
) IShopService {
	return &ShopService{
		ShopRepository:        shopRepository,
		CartRepository:        cartRepository,
		AclRoleUserRepository: aclRoleUserRepository,
	}
}

func (s *ShopService) CreateShopForNewAccount(request *model.CreateShopForNewUserRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		ownerId := request.UserID
		shopName := request.ShopName
		if shopName == "" {
			shopName = "New Shop " + strconv.FormatInt(ownerId, 10)
		}
		_, err := s.ShopRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Shop.FkUser,
				table.Shop.Name,
			},
			entity.Shop{
				FkUser: ownerId,
				Name:   shopName,
			},
		)
		if err != nil {
			return nil, err
		}
		_, err = s.CartRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Cart.FkUser,
			},
			entity.Cart{
				FkUser: ownerId,
			},
		)
		return nil, err
	}

	_, err := s.ShopRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShopService) GetShop(shopId int64) (*entity.Shop, error) {
	return s.ShopRepository.GetById(s.ShopRepository.GetDatabase().Db, shopId)
}

func (s *ShopService) UpdateShop(shopId int64, request *model.UpdateShopRequest) error {
	updatedColumnList := postgres.ColumnList{}
	shop := entity.Shop{
		IDShop: shopId,
	}

	if request.Name != nil {
		updatedColumnList = append(
			updatedColumnList,
			table.Shop.Name,
		)
		shop.Name = *request.Name
	}

	if request.Description != nil {
		updatedColumnList = append(
			updatedColumnList,
			table.Shop.Description,
		)
		shop.Description = request.Description
	}

	if _, err := s.ShopRepository.UpdateById(
		s.ShopRepository.GetDatabase().Db,
		updatedColumnList,
		shop,
	); err != nil {
		return err
	}
	return nil
}

func (s *ShopService) IsFollowed(shopId, userId int64) (bool, error) {
	if _, err := s.AclRoleUserRepository.GetByParam(
		s.AclRoleUserRepository.GetDatabase().Db,
		&entity.ACLRoleUser{
			FkACLRole: constants.FOLLOWER,
			FkUser:    userId,
			FkShop:    shopId,
		},
	); err != nil {
		/*
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code != pgerrcode.NoDataFound {
					return false, nil
				}
			}
			return false, err*/
		if err.Error() == "qrm: no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *ShopService) AddFollower(shopId, userId int64) error {
	if _, err := s.AclRoleUserRepository.CreateOne(
		s.AclRoleUserRepository.GetDatabase().Db,
		postgres.ColumnList{
			table.ACLRoleUser.FkACLRole,
			table.ACLRoleUser.FkUser,
			table.ACLRoleUser.FkShop,
		},
		entity.ACLRoleUser{
			FkACLRole: constants.FOLLOWER,
			FkUser:    userId,
			FkShop:    shopId,
		},
	); err != nil {
		return err
	}

	return nil
}
