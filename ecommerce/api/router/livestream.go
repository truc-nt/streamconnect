package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadLivestreamRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	livestream := apiRouter.Group("/livestreams")
	{
		livestream.GET("/", h.LivestreamHandler.GetLivestreams)
		livestream.GET("/:livestream_id", h.LivestreamHandler.GetLivestream)
		livestream.GET("/:livestream_id/info", h.LivestreamHandler.GetLivestreamInfo)
		livestream.GET("/:livestream_id/livestream_products", h.LivestreamProductHandler.GetLivestreamProductsByLivestreamId)

		livestream.Use(AuthorizationMiddleware())

		livestream.POST("/:livestream_id/save_hls", h.LivestreamHandler.SetLivestreamHls)
		livestream.POST("/:livestream_id/add_livestream_product", h.LivestreamHandler.AddLivestreamProduct)
		livestream.PUT("/:livestream_id/start_livestream", h.LivestreamHandler.StartLivestream)
		livestream.POST("/:livestream_id/register_product_follower", h.LivestreamHandler.RegisterLivestreamProductFollower)
		livestream.GET("/-/livestream_products/:livestream_product_id/followers",
			h.LivestreamHandler.FetchLivestreamProductFollowers,
		)
	}
}
