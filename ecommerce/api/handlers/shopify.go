package handlers

import (
	"ecommerce/api/models"
	"ecommerce/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IShopifyHandler interface {
	Connect(*gin.Context)
	Redirect(*gin.Context)
}

type ShopifyHandler struct {
	BaseHandler
	Service services.IShopifyService
}

func NewShopifyHandler(s services.IShopifyService) IShopifyHandler {
	return &ShopifyHandler{
		Service: s,
	}
}

func (h *ShopifyHandler) Connect(ctx *gin.Context) {
	var queryParams *models.ShopifyConnectParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	authorizePath := h.Service.GetAuthorizePath(queryParams.Shop)

	ctx.Redirect(http.StatusMovedPermanently, authorizePath)
}

func (h *ShopifyHandler) Redirect(ctx *gin.Context) {
	var queryParams *models.ShopifyRedirectParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.ConnectNewShopifyExternalShop(queryParams.Shop, queryParams.Code); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
}
