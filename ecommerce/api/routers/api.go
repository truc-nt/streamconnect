package routers

import (
	"ecommerce/api/handlers"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(e *gin.Engine, h *handlers.Handlers) {
	apiRouter := e.Group("/api")
	{
		LoadShopifyRouter(apiRouter, h.ShopifyHandler)
	}
}
