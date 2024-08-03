package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadProductRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	externalShopRouter := apiRouter.Group("/products")
	{
		externalShopRouter.POST("/", h.ProductHandler.CreateProductsVariants)

	}
}
