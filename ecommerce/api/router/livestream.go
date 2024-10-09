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
		livestream.GET("/:livestream_id/livestream_products", h.LivestreamProductHandler.GetLivestreamProductsByLivestreamId)

		livestream.Use(AuthorizationMiddleware())

		livestream.POST("/:livestream_id/save_hls", h.LivestreamHandler.SetLivestreamHls)
		livestream.POST("/:livestream_id/add_livestream_product", h.LivestreamHandler.AddLivestreamProduct)
		livestream.PATCH("/:livestream_id", h.LivestreamHandler.UpdateLivestream)
		livestream.POST("/:livestream_id/livestream_products/follow", h.LivestreamHandler.RegisterLivestreamProductFollower)
		livestream.GET("/:livestream_id/livestream_products/follow", h.LivestreamHandler.GetFollowLivestreamProductsInLivestream)
		livestream.GET("/:livestream_id/statistics", h.LivestreamHandler.GetLivestreamStatistics)

		livestream.GET("/-/livestream_products/:livestream_product_id/followers",
			h.LivestreamHandler.FetchLivestreamProductFollowers,
		)

		livestream.PATCH("/:livestream_id/update_livestream_products", h.LivestreamHandler.UpdateLivestreamProducts)
	}
}
