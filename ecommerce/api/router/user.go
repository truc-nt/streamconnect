package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadUserRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	userRouter := apiRouter.Group("/users")
	{
		userRouter.GET("/:user_id/address", h.UserHandler.GetDefaultAddress)
	}
}
