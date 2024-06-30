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
	SyncProducts(*gin.Context)

	Authorize(*gin.Context)
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

func (h *ShopifyHandler) Authorize(ctx *gin.Context) {
	authorizePath := h.Service.GetAuthorizePath(2, "storetest306", "3d28749d60b6b97e318c24185a0fd2b3", "6ffa4b351c128862486c337595249835")
	ctx.Redirect(http.StatusMovedPermanently, authorizePath)
}

func (h *ShopifyHandler) Connect(ctx *gin.Context) {
	var req *models.ShopifyConnectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.HandleFailed(ctx, err)
		return
	}

	if err := h.Service.CreateNewShopifyAuth(req.UserId, req.ShopName, req.ClientId, req.ClientSecret); err != nil {
		h.HandleFailed(ctx, err)
		return
	}

	//h.HandleSuccessCreate(ctx)

	authorizePath := h.Service.GetAuthorizePath(req.UserId, req.ShopName, req.ClientId, req.ClientSecret)
	ctx.Redirect(http.StatusMovedPermanently, authorizePath)
}

func (h *ShopifyHandler) Redirect(ctx *gin.Context) {
	var queryParams *models.ShopifyRedirectParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.HandleFailed(ctx, err)
		return
	}

	if err := h.Service.SaveAccessToken(queryParams.UserId, queryParams.Code); err != nil {
		h.HandleFailed(ctx, err)
		return
	}
	h.HandleSuccessCreate(ctx)
}

func (h *ShopifyHandler) SyncProducts(ctx *gin.Context) {
	var queryParams *models.ShopifySyncProductsParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.HandleFailed(ctx, err)
		return
	}

	if err := h.Service.SyncProducts(queryParams.UserId); err != nil {
		h.HandleFailed(ctx, err)
		return
	}

	h.HandleSuccessUpdate(ctx)
}
