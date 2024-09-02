package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadExternalProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	externalProductRouter := apiRouter.Group("/external_products")
	{
		externalProductRouter.GET("/", h.ExternalVariantHandler.GetExternalVariantsGroupByProduct)
		externalProductRouter.GET("/:external_product_id_mapping", h.ExternalVariantHandler.GetExternalVariantsByExternalProductIdMapping)
		externalProductRouter.GET("/:external_product_id_mapping/variants", h.VariantHandler.GetVariantsByExternalProductIdMapping)
	}

	externalVariantRouter := apiRouter.Group("/external_variants")
	{
		externalVariantRouter.POST("/connect", h.ExternalVariantHandler.ConnectVariants)
	}
}
