package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(e *gin.Engine, h *handler.Handlers) {
	apiRouter := e.Group("/api")
	{
		LoadProductRouter(apiRouter, h)
		LoadShopRouter(apiRouter, h)
		LoadExternalShopRouter(apiRouter, h)
		LoadShopifyRouter(apiRouter, h)
		LoadLivestreamRouter(apiRouter, h)
		LoadLivestreamProductRouter(apiRouter, h)
		LoadCartRouter(apiRouter, h)
	}
}
