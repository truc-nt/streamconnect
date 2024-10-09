package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadOrderRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	orderRouter := apiRouter.Group("/orders")
	orderRouter.Use(AuthorizationMiddleware())
	{
		orderRouter.GET("/buy", h.OrderHandler.GetBuyOrders)
		orderRouter.GET("/:order_id", h.OrderHandler.GetOrder)
		orderRouter.POST("/create_with_cart_items", h.OrderHandler.CreateOrderWithCartItems)
		orderRouter.POST("/livestream_ext_variants/:livestream_ext_variant_id", h.OrderHandler.CreateOrderWithLivestreamExtVariantId)
	}
}
