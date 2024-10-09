package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	productRouter := apiRouter.Group("/products")
	{
		productRouter.GET("/:product_id", h.ProductHandler.GetProductById)
		productRouter.GET("/:product_id/variants", h.VariantHandler.GetVariantsByProductId)
		productRouter.PATCH("/:product_id", h.ProductHandler.UpdateProduct)
	}

	variantRouter := apiRouter.Group("/variants")
	{
		variantRouter.GET("/:variant_id/external_variants", h.VariantHandler.GetExternalVariantsByVariantId)
	}
}
