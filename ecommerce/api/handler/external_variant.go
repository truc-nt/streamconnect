package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IExternalVariantHandler interface {
	GetExternalVariantsGroupByProduct(ctx *gin.Context)
	GetExternalVariantsByExternalProductIdMapping(ctx *gin.Context)

	ConnectVariants(ctx *gin.Context)
}

type ExternalVariantHandler struct {
	BaseHandler
	Service service.IExternalVariantService
}

func NewExternalVariantHandler(s service.IExternalVariantService) IExternalVariantHandler {
	return &ExternalVariantHandler{
		Service: s,
	}
}

func (h *ExternalVariantHandler) GetExternalVariantsGroupByProduct(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParams *model.GetExternalVariantsGroupByProduct
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	products, err := h.Service.GetExternalVariantsGroupByProduct(shopId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, products)
}

func (h *ExternalVariantHandler) GetExternalVariantsByExternalProductIdMapping(ctx *gin.Context) {
	externalVariants, err := h.Service.GetExternalVariantsByExternalProductIdMapping(ctx.Param("external_product_id_mapping"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, externalVariants)
}

func (h *ExternalVariantHandler) ConnectVariants(ctx *gin.Context) {
	var connectVariantsRequest *model.ConnectVariantsRequest
	if err := ctx.ShouldBindJSON(&connectVariantsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.ConnectVariants(connectVariantsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessUpdate(ctx)
}
