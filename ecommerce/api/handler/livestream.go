package handler

import (
	"database/sql"
	"ecommerce/api/model"
	"ecommerce/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
)

type ILivestreamHandler interface {
	CreateLivestream(ctx *gin.Context)
	FetchLivestreams(ctx *gin.Context)
	GetLivestream(ctx *gin.Context)
	SetLivestreamHls(ctx *gin.Context)
	RegisterLivestreamProductFollower(ctx *gin.Context)
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

func (h *LivestreamHandler) GetLivestream(ctx *gin.Context) {
	id, err := h.parseId(ctx, ctx.Param("livestream_id"))
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

func (h *LivestreamHandler) FetchLivestreams(ctx *gin.Context) {
	status := ctx.Query("status")
	nillAbleStatus := sql.NullString{
		String: status,
		Valid:  status != "",
	}
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	nillAbleShopId := sql.NullInt64{
		Int64: shopId,
		Valid: err == nil,
	}

	livestreams, err := h.Service.FetchLivestreams(nillAbleStatus, nillAbleShopId)
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
	if request.IDLivestream != id {
		h.handleFailed(ctx, errors.New("bad request"))
	}
	if err := h.Service.SetLivestreamHls(request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}

func (h *LivestreamHandler) RegisterLivestreamProductFollower(ctx *gin.Context) {
	idLivestream, err := h.parseId(ctx, ctx.Param("livestream_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	var request *model.RegisterLivestreamProductFollowerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	if request.IDLivestream != idLivestream {
		h.handleFailed(ctx, errors.New("livestream id not match"))
		return
	}
	if err := h.Service.RegisterLivestreamProductFollower(request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
}
