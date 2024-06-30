package routers

import (
	"ecommerce/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadShopifyRouter(apiRouter *gin.RouterGroup, shopifyHandler handlers.IShopifyHandler) {
	shopifyRouter := apiRouter.Group("/shopify")
	{
		shopifyRouter.POST("/connect", shopifyHandler.Connect)
		shopifyRouter.GET("/redirect", shopifyHandler.Redirect)
		shopifyRouter.GET("/sync-products", shopifyHandler.SyncProducts)

		shopifyRouter.GET("/authorize", shopifyHandler.Authorize)
	}
}
