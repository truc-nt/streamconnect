package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	GetProductById(ctx *gin.Context)
	GetProductsByShopId(ctx *gin.Context)
	CreateProductsWithVariants(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
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

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	productId, err := h.parseId(ctx, ctx.Param("product_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	product, err := h.Service.GetProductById(productId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, product)
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

	h.handleSuccessGet(ctx, products)
}

func (h *ProductHandler) CreateProductsWithVariants(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var createProductsWithVariantsRequest *model.CreateProductWithVariants
	if err := ctx.ShouldBindJSON(&createProductsWithVariantsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.CreateProductsWithVariants(shopId, createProductsWithVariantsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	productId, err := h.parseId(ctx, ctx.Param("product_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var updateProductRequest *model.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&updateProductRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.UpdateProduct(productId, updateProductRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}
