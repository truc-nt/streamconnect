package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadOrderRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	orderRouter := apiRouter.Group("/orders")
	{
		orderRouter.POST("/create_with_cart_items", h.OrderHandler.CreateOrderWithCartItems)
	}
}
