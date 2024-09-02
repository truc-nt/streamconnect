package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadExternalShopRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	externalShopRouter := apiRouter.Group("/external_shops")
	{
		externalShopRouter.POST("/:external_shop_id/sync_external_variants", h.ExternalShopHandler.SyncExternalVariantsByExternalShopId)
	}
}
