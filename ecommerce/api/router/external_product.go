package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadExternalProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	externalVariantRouter := apiRouter.Group("/external_products")
	{
		externalVariantRouter.GET("/", h.ExternalProductHandler.GetExternalVariantsGroupByProduct)
		externalVariantRouter.GET("/:external_product_id_mapping", h.ExternalProductHandler.GetExternalVariantsByExternalProductIdMapping)
	}
}
