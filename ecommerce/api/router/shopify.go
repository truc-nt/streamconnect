package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadShopifyRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	shopifyRouter := apiRouter.Group("/shopify")
	{
		shopifyRouter.GET("/connect", h.ShopifyHandler.Connect)
		shopifyRouter.GET("/redirect", h.ShopifyHandler.Redirect)
	}
}
