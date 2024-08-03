package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IExternalShopHandler interface {
	GetExternalShopsByShopId(ctx *gin.Context)
	SyncExternalShopsByExternalShopId(ctx *gin.Context)
}

type ExternalShopHandler struct {
	BaseHandler
	Service service.IExternalShopService
}

func NewExternalShopHandler(s service.IExternalShopService) IExternalShopHandler {
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

	var queryParams *model.GetExternalShopsByShopIdParam
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

func (h *ExternalShopHandler) SyncExternalShopsByExternalShopId(ctx *gin.Context) {
	externalShopId, err := h.parseId(ctx, ctx.Param("external_shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.SyncExternalProductsByExternalShopId(externalShopId); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}
