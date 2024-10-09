package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadLivestreamProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	livestreamProduct := apiRouter.Group("/livestream_products")
	{
		livestreamProduct.GET("/:livestream_product_id", h.LivestreamProductHandler.GetLivestreamProductInfoByLivestreamProductId)
		livestreamProduct.POST("/priority", h.LivestreamProductHandler.UpdateLivestreamProductPriority)

		livestreamProduct.DELETE("/:livestream_product_id/follow", h.LivestreamHandler.DeleteLivestreamProductFollower)
	}

	livestreamExternalVariant := apiRouter.Group("/livestream_external_variants")
	{
		livestreamExternalVariant.POST("/update_quantity", h.LivestreamHandler.UpdateLivestreamExternalVariantQuantity)

	}
}
