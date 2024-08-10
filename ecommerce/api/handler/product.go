package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	GetProductsByShopId(ctx *gin.Context)
	CreateProductsVariants(ctx *gin.Context)
}

type ProductHandler struct {
	BaseHandler
	Service service.IProductService
}

func NewProductHandler(s service.IProductService) IProductHandler {
	return &ProductHandler{
		Service: s,
	}
}

func (h *ProductHandler) GetProductsByShopId(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParam *model.GetProductsByShopIdParam
	if err := ctx.ShouldBindQuery(&queryParam); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	products, err := h.Service.GetProductsByShopId(shopId, queryParam.Limit, queryParam.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	/*var res []*model.GetProductsByShopIdData
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &res,
	})
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := decoder.Decode(products); err != nil {
		h.handleFailed(ctx, err)
		return
	}*/

	h.handleSuccessGet(ctx, products)
}

func (h *ProductHandler) CreateProductsVariants(ctx *gin.Context) {
	var createProductsRequest *model.CreateProductsVariantsRequest
	if err := ctx.ShouldBindJSON(&createProductsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.CreateProductsVariantsFromExternalProducts(createProductsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}
