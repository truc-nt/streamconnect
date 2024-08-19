package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadCartRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	cart := apiRouter.Group("/cart")
	{
		cart.GET("/:cart_id", h.CartHandler.Get)
		cart.POST("/:cart_id", h.CartHandler.AddToCart)
	}
}
