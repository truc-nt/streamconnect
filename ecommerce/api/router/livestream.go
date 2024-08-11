package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadLivestreamRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	livestream := apiRouter.Group("/livestreams")
	{
		livestream.GET("/:livestream_id/products", h.LivestreamHandler.GetProductsByLivestreamId)
	}
}
