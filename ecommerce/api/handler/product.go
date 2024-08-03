package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
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
