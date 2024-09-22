package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadApiRouter(e *gin.Engine, h *handler.Handlers) {
	apiRouter := e.Group("/api")
	{
		LoadUserRouter(apiRouter, h)
		LoadProductRouter(apiRouter, h)
		LoadExternalProductRouter(apiRouter, h)
		LoadShopRouter(apiRouter, h)
		LoadExternalShopRouter(apiRouter, h)
		LoadShopifyRouter(apiRouter, h)
		LoadLivestreamRouter(apiRouter, h)
		LoadLivestreamProductRouter(apiRouter, h)
		LoadCartRouter(apiRouter, h)
		LoadOrderRouter(apiRouter, h)
		LoadVoucherRouter(apiRouter, h)
	}
}
