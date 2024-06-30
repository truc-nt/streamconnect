package shopify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"ecommerce/internal/constants"

	"golang.org/x/oauth2"
)

//go:generate mockgen -package=shopify -source=shopify_client.go -destination=shopify_client_mock.go
type IShopifyClient interface {
	GetAuthorizePath(oauth2Config *oauth2.Config) string
	GetAccessToken(oauth2Config *oauth2.Config, code string) (*oauth2.Token, error)
	GetProducts() (*GetProductsResponse, error)
}

type ShopifyClientParam struct {
	ShopName    string
	AccessToken string
}

type ShopifyClient struct {
	Client *http.Client

	Param *ShopifyClientParam
}

// Create new Shopify client
func NewShopifyClient(param *ShopifyClientParam) IShopifyClient {
	c := &ShopifyClient{
		Client: &http.Client{},
		Param:  param,
	}

	return c
}

func (c *ShopifyClient) GetAuthorizePath(oauth2Config *oauth2.Config) string {
	oauth2Config.Endpoint = oauth2.Endpoint{
		AuthURL:  fmt.Sprintf(constants.ShopifyBaseURL, c.Param.ShopName) + constants.ShopifyAuthorizePath,
		TokenURL: fmt.Sprintf(constants.ShopifyBaseURL, c.Param.ShopName) + constants.ShopifyAccessTokenPath,
	}
	return oauth2Config.AuthCodeURL("state")
}

func (c *ShopifyClient) GetAccessToken(oauth2Config *oauth2.Config, code string) (*oauth2.Token, error) {
	oauth2Config.Endpoint = oauth2.Endpoint{
		AuthURL:  fmt.Sprintf(constants.ShopifyBaseURL, c.Param.ShopName) + constants.ShopifyAuthorizePath,
		TokenURL: fmt.Sprintf(constants.ShopifyBaseURL, c.Param.ShopName) + constants.ShopifyAccessTokenPath,
	}

	res, err := oauth2Config.Exchange(context.Background(), code)
	return res, err
}

func (c *ShopifyClient) GetProducts() (*GetProductsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", fmt.Sprintf(constants.ShopifyBaseURL, c.Param.ShopName), constants.ShopifyGetProductsPath), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constants.ShopifyTokenKey, c.Param.AccessToken)

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
