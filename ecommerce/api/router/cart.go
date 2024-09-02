package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadCartRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	cart := apiRouter.Group("/carts")
	{
		cart.GET("/:cart_id", h.CartHandler.Get)
		cart.POST("/:cart_id", h.CartHandler.AddToCart)
	}

	cartItem := apiRouter.Group("/cart_items")
	{
		cartItem.PATCH("/:cart_item_id", h.CartHandler.Update)
		cartItem.POST("/get_cart_items_by_ids", h.CartHandler.GetCartItemsByIds)
	}
}
