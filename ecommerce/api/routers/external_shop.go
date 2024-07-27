package routers

import (
	"ecommerce/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadExternalShopRouter(apiRouter *gin.RouterGroup, externalShopHandler handlers.IExternalShopHandler) {
	externalShopRouter := apiRouter.Group("/external_shops")
	{
		externalShopRouter.GET("/:shop_id", externalShopHandler.GetExternalShopsByShopId)
		externalShopRouter.GET("/sync_products/:external_shop_id", externalShopHandler.SyncProductsByExternalShopId)

	}
}
