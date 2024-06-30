package services

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/configs"
	"ecommerce/internal/models"
	"ecommerce/internal/repositories"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"

	"github.com/jackc/pgx"
)

type IShopifyService interface {
	//GetCredential(userId int) (models.UserShopifyAuth, err error)
	CreateNewShopifyAuth(userId int32, shopName string, clientId string, clientSecret string) error
	GetAuthorizePath(userId int32, shopName string, clientId string, clientSecret string) string
	SaveAccessToken(userId int32, code string) error
	SyncProducts(userId int32) error
}

type ShopifyService struct {
	Config     *configs.Config
	Repository repositories.IShopifyRepository
}

func NewShopifyService(repo repositories.IShopifyRepository, config *configs.Config) IShopifyService {
	return &ShopifyService{
		Repository: repo,
		Config:     config,
	}
}

//func (s *ShopifyService) GetCredential(userId int) (clientId string, clientSecret string, accessToken string, err error) {

func (s *ShopifyService) CreateNewShopifyAuth(userId int32, shopName string, clientId string, clientSecret string) error {
	data := models.UserShopifyAuth{
		FkUser:       userId,
		ShopName:     shopName,
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
	if err := s.Repository.Create(data); err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("user already connected to this Shopify store")
			}
		}
		return fmt.Errorf("failed to create new Shopify auth: %v", err)
	}
	return nil
}

func (s *ShopifyService) GetAuthorizePath(userId int32, shopName string, clientId string, clientSecret string) string {
	adapter := adapters.NewShopifyAdapter(&adapters.ShopifyAdapterConfig{
		ShopName:     shopName,
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}, s.Config)
	return adapter.GetAuthorizePath(userId)
}

func (s *ShopifyService) SaveAccessToken(userId int32, code string) error {
	data, err := s.Repository.GetByUserId(int(userId))
	if err != nil {
		return err
	}

	adapter := adapters.NewShopifyAdapter(&adapters.ShopifyAdapterConfig{
		ShopName:     data.ShopName,
		ClientID:     data.ClientID,
		ClientSecret: data.ClientSecret,
	}, s.Config)

	accessToken, err := adapter.GetAccessToken(code)
	if err != nil {
		return err
	}

	data.AccessToken = accessToken
	if err := s.Repository.UpdateByUserId(int(userId), data); err != nil {
		return fmt.Errorf("failed to save access token: %v", err)
	}
	return nil
}

func (s *ShopifyService) SyncProducts(userId int32) error {
	data, err := s.Repository.GetByUserId(int(userId))
	if err != nil {
		return err
	}

	fmt.Printf("%v", data)

	adapter := adapters.NewShopifyAdapter(&adapters.ShopifyAdapterConfig{
		ShopName:     data.ShopName,
		ClientID:     data.ClientID,
		ClientSecret: data.ClientSecret,
		AccessToken:  data.AccessToken,
	}, s.Config)

	_, err = adapter.GetProducts()
	if err != nil {
		return err
	}
	return nil
}
