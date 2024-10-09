package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadShopRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	shopRouter := apiRouter.Group("/shops")
	shopRouter.GET("/:shop_id", h.ShopHandler.GetShop)
	shopRouter.Use(AuthorizationMiddleware())
	{
		shopRouter.PATCH("/:shop_id", h.ShopHandler.UpdateShop)
		shopRouter.POST("/:shop_id/follow", h.ShopHandler.FollowShop)
		shopRouter.GET("/:shop_id/products", h.ProductHandler.GetProductsByShopId)
		shopRouter.GET("/:shop_id/orders", h.OrderHandler.GetOrdersByShopId)

		shopRouter.POST("/:shop_id/products/", h.ProductHandler.CreateProductsWithVariants)
		shopRouter.GET("/:shop_id/external_products", h.ExternalVariantHandler.GetExternalVariantsGroupByProduct)
		shopRouter.GET("/:shop_id/external_shops", h.ExternalShopHandler.GetExternalShopsByShopId)
		shopRouter.POST("/:shop_id/livestreams/create", h.LivestreamHandler.CreateLivestream)
		shopRouter.POST("/forNewUser", h.ShopHandler.CreateShopForNewAccount)

		shopRouter.GET("/:shop_id/vouchers", h.VoucherHandler.GetShopVouchers)
	}
}
