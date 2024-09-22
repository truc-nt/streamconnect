package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadUserRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	userRouter := apiRouter.Group("/users")
	userRouter.Use(AuthorizationMiddleware())
	{
		userRouter.GET("/", h.UserHandler.GetUser)

	}

	addressRouter := apiRouter.Group("/addresses")
	addressRouter.Use(AuthorizationMiddleware())
	{
		addressRouter.GET("/default_address", h.UserHandler.GetDefaultAddress)
		addressRouter.GET("/", h.UserHandler.GetAddressesByUserId)
		addressRouter.POST("/", h.UserHandler.CreateAddress)
	}
}
