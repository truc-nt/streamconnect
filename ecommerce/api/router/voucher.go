package router

import (
	"ecommerce/api/handler"

	"github.com/gin-gonic/gin"
)

func LoadVoucherRouter(apiRouter *gin.RouterGroup, h *handler.Handlers) {
	vouchersRoute := apiRouter.Group("/vouchers")
	vouchersRoute.Use(AuthorizationMiddleware())
	{
		shopVouchersRoute := vouchersRoute.Group("/shop")
		{
			//shopVouchersRoute.GET("/", h.VoucherHandler.GetShopVouchers)
			//shopVouchersRoute.GET("/:shop_id/user", h.VoucherHandler.GetShopVouchersWithUserId)
			shopVouchersRoute.POST("/", h.VoucherHandler.Create)
		}

		userVouchersRoute := vouchersRoute.Group("/user")
		{
			userVouchersRoute.PUT("/:voucher_id", h.VoucherHandler.AddVoucher)
		}
	}
}
