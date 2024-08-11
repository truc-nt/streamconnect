package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IVariantHandler interface {
	GetVariantsByProductId(ctx *gin.Context)
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

	/*product, err := h.ProductService.GetProductById(productId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}*/

	variants, err := h.VariantService.GetVariantsByProductId(productId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	/*res := &model.GetVariantsByProductIdData{
		Name:   product.Name,
		Status: product.Status,
		Option: product.OptionTitles,
	}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &res.Variants,
	})

	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := decoder.Decode(variants); err != nil {
		h.handleFailed(ctx, err)
		return
	}*/

	h.handleSuccessGet(ctx, variants)
}
