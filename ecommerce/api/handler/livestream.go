package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ILivestreamHandler interface {
	CreateLivestream(ctx *gin.Context)
}

type LivestreamHandler struct {
	BaseHandler
	Service service.ILivestreamService
}

func NewLivestreamHandler(s service.ILivestreamService) ILivestreamHandler {
	return &LivestreamHandler{
		Service: s,
	}
}

func (h *LivestreamHandler) CreateLivestream(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var createLivestreamRequest *model.CreateLivestreamRequest
	if err := ctx.ShouldBindJSON(&createLivestreamRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.CreateLivestream(shopId, createLivestreamRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}
