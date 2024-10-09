package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IVariantHandler interface {
	GetVariantsByProductId(ctx *gin.Context)
	GetVariantsByExternalProductIdMapping(ctx *gin.Context)
	GetExternalVariantsByVariantId(ctx *gin.Context)
}

type VariantHandler struct {
	BaseHandler
	VariantService service.IVariantService
	ProductService service.IProductService
}

func NewVariantHandler(productService service.IProductService, variantService service.IVariantService) IVariantHandler {
	return &VariantHandler{
		VariantService: variantService,
		ProductService: productService,
	}
}

func (h *VariantHandler) GetVariantsByProductId(ctx *gin.Context) {
	productId, err := h.parseId(ctx, ctx.Param("product_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParams *model.GetVariantsByProductIdParam
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	variants, err := h.VariantService.GetVariantsByProductId(productId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, variants)
}

func (h *VariantHandler) GetVariantsByExternalProductIdMapping(ctx *gin.Context) {
	externalProductIdMapping := ctx.Param("external_product_id_mapping")
	variants, err := h.VariantService.GetVariantsByExternalProductIdMapping(externalProductIdMapping)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, variants)
}

func (h *VariantHandler) GetExternalVariantsByVariantId(ctx *gin.Context) {
	variantId, err := h.parseId(ctx, ctx.Param("variant_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	externalVariants, err := h.VariantService.GetExternalVariantsByVariantId(variantId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, externalVariants)
}
