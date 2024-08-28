package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadShopRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	shopRouter := apiRouter.Group("/shops")
	{
		shopRouter.GET("/:shop_id/products", h.ProductHandler.GetProductsByShopId)
		shopRouter.POST("/:shop_id/products/", h.ProductHandler.CreateProductWithVariants)
		shopRouter.GET("/:shop_id/external_shops", h.ExternalShopHandler.GetExternalShopsByShopId)
		shopRouter.POST("/:shop_id/livestreams/create", h.LivestreamHandler.CreateLivestream)
	}
}
