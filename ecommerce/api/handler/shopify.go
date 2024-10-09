package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IShopifyHandler interface {
	Connect(*gin.Context)
	Redirect(*gin.Context)
}

type ShopifyHandler struct {
	BaseHandler
	Service service.IShopifyService
}

func NewShopifyHandler(s service.IShopifyService) IShopifyHandler {
	return &ShopifyHandler{
		Service: s,
	}
}

func (h *ShopifyHandler) Connect(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParams *model.ShopifyConnectParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	authorizePath := h.Service.GetAuthorizePath(shopId, queryParams.Shop)

	//ctx.Redirect(http.StatusMovedPermanently, authorizePath)
	h.handleSuccessGet(ctx, authorizePath)
}

func (h *ShopifyHandler) Redirect(ctx *gin.Context) {

	var queryParams *model.ShopifyRedirectParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.ConnectNewExternalShopShopify(queryParams.State, queryParams.Shop, queryParams.Code); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
	ctx.Redirect(http.StatusMovedPermanently, "http://localhost:3000/seller/shops")
}
