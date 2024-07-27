package handlers

import (
	"ecommerce/api/models"
	"ecommerce/internal/services"

	"github.com/gin-gonic/gin"
)

type IExternalShopHandler interface {
	GetExternalShopsByShopId(ctx *gin.Context)
	SyncProductsByExternalShopId(ctx *gin.Context)
}

type ExternalShopHandler struct {
	BaseHandler
	Service services.IExternalShopService
}

func NewExternalShopHandler(s services.IExternalShopService) IExternalShopHandler {
	return &ExternalShopHandler{
		Service: s,
	}
}

func (h *ExternalShopHandler) GetExternalShopsByShopId(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParams *models.GetExternalShopsByShopIdParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	shops, err := h.Service.GetExternalShopsByShopId(shopId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, shops)
}

func (h *ExternalShopHandler) SyncProductsByExternalShopId(ctx *gin.Context) {
	externalShopId, err := h.parseId(ctx, ctx.Param("external_shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.SyncProductsByExternalShopId(externalShopId); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}
