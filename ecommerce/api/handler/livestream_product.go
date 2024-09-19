package handler

import (
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ILivestreamProductHandler interface {
	GetLivestreamProductsByLivestreamId(ctx *gin.Context)
	GetLivestreamProductInfoByLivestreamProductId(ctx *gin.Context)
	FetchLivestreamProductFollowers(ctx *gin.Context)
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

func (h *LivestreamProductHandler) GetLivestreamProductsByLivestreamId(ctx *gin.Context) {
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

func (h *LivestreamProductHandler) FetchLivestreamProductFollowers(ctx *gin.Context) {
	id, err := h.parseId(ctx, ctx.Param("livestream_product_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	followers, err := h.Service.FetchLivestreamProductFollowers(id)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, followers)
}
