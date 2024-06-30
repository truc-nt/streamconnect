package adapters

import (
	"ecommerce/internal/clients/shopify"
	"ecommerce/internal/configs"
	"fmt"
)

const (
	ShopifyBaseURL         = "https://%s.myshopify.com"
	ShopifyTokenKey        = "X-Shopify-Access-Token"
	ShopifyAuthorizePath   = "/admin/oauth/authorize"
	ShopifyRedirectPath    = "/api/shopify/redirect"
	ShopifyAccessTokenPath = "/admin/oauth/access_token"
)

type IShopifyAdapter interface {
	GetShopifyClient() shopify.IShopifyClient
	GetAuthorizePath(userId int32) string
	GetAccessToken(code string) (string, error)
	GetProducts() (interface{}, error)
	//GetShopifyAuthConfig() *oauth2.Config
}

type ShopifyAdapterConfig struct {
	ShopName     string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

type ShopifyAdapter struct {
	Config    *ShopifyAdapterConfig
	AppConfig *configs.Config
}

func NewShopifyAdapter(cfg *ShopifyAdapterConfig, appCfg *configs.Config) IShopifyAdapter {
	return &ShopifyAdapter{
		Config:    cfg,
		AppConfig: appCfg,
	}
}

func (a *ShopifyAdapter) GetShopifyClient() shopify.IShopifyClient {
	param := &shopify.ShopifyClientConfig{
		ShopName:     a.Config.ShopName,
		ClientID:     a.Config.ClientID,
		ClientSecret: a.Config.ClientSecret,
		AccessToken:  a.Config.AccessToken,
		HostBaseUrl:  fmt.Sprintf("%s:%d", a.AppConfig.Server.Host, a.AppConfig.Server.Port),
	}

	return shopify.NewShopifyClient(param)
}

func (a *ShopifyAdapter) GetAuthorizePath(userId int32) string {
	return a.GetShopifyClient().GetAuthorizePath(userId)
}

func (a *ShopifyAdapter) GetAccessToken(code string) (string, error) {
	token, err := a.GetShopifyClient().GetAccessToken(code)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (a *ShopifyAdapter) GetProducts() (interface{}, error) {
	products, err := a.GetShopifyClient().GetProducts()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", products)
	return products, nil
}
