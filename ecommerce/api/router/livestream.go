package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadLivestreamRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	livestream := apiRouter.Group("/livestreams")
	{
		livestream.GET("/", h.LivestreamHandler.FetchLivestreams)
		livestream.GET("/:livestream_id", h.LivestreamHandler.GetLivestream)
		livestream.GET("/:livestream_id/livestream_products", h.LivestreamProductHandler.GetLivestreamProductsByLivestreamId)
		livestream.POST("/:livestream_id/startHls", h.LivestreamHandler.SetLivestreamHls)
	}
}
