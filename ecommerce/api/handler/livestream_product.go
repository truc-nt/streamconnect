package handler

import (
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ILivestreamProductHandler interface {
	GetProductsByLivestreamId(ctx *gin.Context)
	GetLivestreamProductInfoByLivestreamProductId(ctx *gin.Context)
}

type LivestreamProductHandler struct {
	BaseHandler
	Service service.ILivestreamProductService
}

func NewLivestreamProductHandler(s service.ILivestreamProductService) ILivestreamProductHandler {
	return &LivestreamProductHandler{
		Service: s,
	}
}

func (h *LivestreamProductHandler) GetProductsByLivestreamId(ctx *gin.Context) {
	livestreamId, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	products, err := h.Service.GetLivestreamProductsByLivestreamId(livestreamId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, products)
}

func (h *LivestreamProductHandler) GetLivestreamProductInfoByLivestreamProductId(ctx *gin.Context) {
	livestreamProductId, err := h.parseId(ctx, ctx.Param("livestream_product_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	product, err := h.Service.GetLivestreamProductInfoByLivestreamProductId(livestreamProductId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, product)
}
