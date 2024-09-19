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

		livestream.Use(AuthorizationMiddleware())
		{
			livestream.GET("/:livestream_id/livestream_products", h.LivestreamProductHandler.GetLivestreamProductsByLivestreamId)
			livestream.POST("/:livestream_id/start_hls", h.LivestreamHandler.SetLivestreamHls)
			livestream.POST("/:livestream_id/register_product_follower", h.LivestreamHandler.RegisterLivestreamProductFollower)
			livestream.GET("/-/livestream_products/:livestream_product_id/followers",
				h.LivestreamHandler.FetchLivestreamProductFollowers,
			)
		}
	}
}
