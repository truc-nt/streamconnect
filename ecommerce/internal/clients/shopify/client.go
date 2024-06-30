package shopify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	ShopifyRedirectPath    = "/api/shopify/redirect"
	ShopifyBaseURL         = "https://%s.myshopify.com"
	ShopifyTokenKey        = "X-Shopify-Access-Token"
	ShopifyAuthorizePath   = "/admin/oauth/authorize"
	ShopifyAccessTokenPath = "/admin/oauth/access_token"
	ShopifyGetProductsPath = "/admin/api/2024-04/products.json"
)

//go:generate mockgen -package=shopify -source=shopify_client.go -destination=shopify_client_mock.go
type IShopifyClient interface {
	GetAuthorizePath(int32) string
	GetAccessToken(string) (*oauth2.Token, error)
	GetProducts() (*GetProductsResponse, error)
}

type ShopifyClientConfig struct {
	ShopName     string
	ClientID     string
	ClientSecret string
	AccessToken  string
	HostBaseUrl  string
}

type ShopifyClient struct {
	Client *http.Client

	Config *ShopifyClientConfig
}

// Create new Shopify client
func NewShopifyClient(clientConfig *ShopifyClientConfig) IShopifyClient {

	c := &ShopifyClient{
		Client: &http.Client{},
		Config: clientConfig,
	}

	return c
}

func (c *ShopifyClient) GetAuthorizePath(userId int32) string {
	oauth2Config := &oauth2.Config{
		ClientID:     c.Config.ClientID,
		ClientSecret: c.Config.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", c.Config.HostBaseUrl, ShopifyRedirectPath),
		Scopes:       []string{"read_products", "read_orders", "write_orders"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf(ShopifyBaseURL, c.Config.ShopName) + ShopifyAuthorizePath,
			TokenURL: fmt.Sprintf(ShopifyBaseURL, c.Config.ShopName) + ShopifyAccessTokenPath,
		},
	}
	return oauth2Config.AuthCodeURL(fmt.Sprintf("%d", userId))
}

func (c *ShopifyClient) GetAccessToken(code string) (*oauth2.Token, error) {
	oauth2Config := &oauth2.Config{
		ClientID:     c.Config.ClientID,
		ClientSecret: c.Config.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s%s", c.Config.HostBaseUrl, ShopifyRedirectPath),
		Scopes:       []string{"read_products", "read_orders", "write_orders"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf(ShopifyBaseURL, c.Config.ShopName) + ShopifyAuthorizePath,
			TokenURL: fmt.Sprintf(ShopifyBaseURL, c.Config.ShopName) + ShopifyAccessTokenPath,
		},
	}

	res, err := oauth2Config.Exchange(context.Background(), code)
	return res, err
}

func (c *ShopifyClient) GetProducts() (*GetProductsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", fmt.Sprintf(ShopifyBaseURL, c.Config.ShopName), ShopifyGetProductsPath), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(ShopifyTokenKey, c.Config.AccessToken)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var products *GetProductsResponse
	if err := json.Unmarshal(resData, &products); err != nil {
		return nil, err
	}

	return products, nil
}
