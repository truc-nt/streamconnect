package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadLivestreamProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	livestreamProduct := apiRouter.Group("/livestream_products")
	{
		livestreamProduct.GET("/:livestream_product_id", h.LivestreamProductHandler.GetLivestreamProductInfoByLivestreamProductId)
		livestreamProduct.POST("/pin", h.LivestreamProductHandler.PinLivestreamProduct)
	}

	livestreamExternalVariant := apiRouter.Group("/livestream_external_variants")
	{
		livestreamExternalVariant.POST("/update_quantity", h.LivestreamHandler.UpdateLivestreamExternalVariantQuantity)

	}
}
