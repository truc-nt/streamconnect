package service

import (
	"ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"strconv"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IShopService interface {
	CreateShopForNewAccount(request *model.CreateShopForNewUserRequest) error
	GetShopInfo(shopId int64) (*entity.Shop, error)
}

type ShopService struct {
	ShopRepository repository.IShopRepository
	CartRepository repository.ICartRepository
}

func NewShopService(shopRepository repository.IShopRepository,
	cartRepository repository.ICartRepository) IShopService {
	return &ShopService{
		ShopRepository: shopRepository,
		CartRepository: cartRepository,
	}
}

func (s *ShopService) CreateShopForNewAccount(request *model.CreateShopForNewUserRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		ownerId := request.UserID
		shopName := request.ShopName
		if shopName == "" {
			shopName = "New Shop - User " + strconv.FormatInt(ownerId, 10)
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

func (s *ShopService) GetShopInfo(shopId int64) (*entity.Shop, error) {
	return s.ShopRepository.GetById(s.ShopRepository.GetDatabase().Db, shopId)
}
