package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ILivestreamHandler interface {
	CreateLivestream(ctx *gin.Context)
	GetLivestreams(ctx *gin.Context)
	GetLivestream(ctx *gin.Context)
	SetLivestreamHls(ctx *gin.Context)
	UpdateLivestreamExternalVariantQuantity(ctx *gin.Context)
	AddLivestreamProduct(ctx *gin.Context)
	StartLivestream(ctx *gin.Context)

	GetLivestreamInfo(ctx *gin.Context)
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

	livestreamId, err := h.Service.CreateLivestream(shopId, createLivestreamRequest)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreateWithData(ctx, livestreamId)
}

func (h *LivestreamHandler) GetLivestream(ctx *gin.Context) {
	id, err := h.parseId(ctx, ctx.Param("livestream_id"))
	//userId := ctx.Param("user_id")
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	livestream, err := h.Service.GetLivestream(id)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, livestream)
}

func (h *LivestreamHandler) GetLivestreams(ctx *gin.Context) {
	var param *model.GetLivestreamsQueryParam
	if err := ctx.ShouldBindQuery(&param); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	livestreams, err := h.Service.GetLivestreams(param)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, livestreams)
}

func (h *LivestreamHandler) SetLivestreamHls(ctx *gin.Context) {
	id, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	var request *model.SetLivestreamHlsRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.SetLivestreamHls(id, request.HlsUrl); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}

func (h *LivestreamHandler) GetLivestreamInfo(ctx *gin.Context) {
	userId := ctx.GetHeader("user_id")

	id, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	livestream, err := h.Service.GetLivestreamInfo(id)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, &model.GetLivestreamInfoResponse{
		MeetingId: livestream.MeetingID,
		IDShop:    livestream.FkShop,
		IsHost:    userId == fmt.Sprintf("%d", livestream.FkShop),
		ShopName:  livestream.ShopName,
	})
}

func (h *LivestreamHandler) UpdateLivestreamExternalVariantQuantity(ctx *gin.Context) {
	var request *model.UpdateLivestreamExternalVariantQuantityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.UpdateLivestreamExternalVariantQuantity(request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}

func (h *LivestreamHandler) AddLivestreamProduct(ctx *gin.Context) {
	var request []*model.LivestreamProductCreateRequest
	livestreamId, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.AddLivestreamProduct(livestreamId, request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}

func (h *LivestreamHandler) StartLivestream(ctx *gin.Context) {
	livestreamId, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.StartLivestream(livestreamId); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessUpdate(ctx)
}
