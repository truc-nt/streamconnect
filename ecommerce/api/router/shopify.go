package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadShopifyRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	shopifyRouter := apiRouter.Group("/shopify")

	shopifyRouter.GET("/redirect", h.ShopifyHandler.Redirect)
	shopifyRouter.Use(AuthorizationMiddleware())
	{
		shopifyRouter.GET("/connect", h.ShopifyHandler.Connect)
	}
}
