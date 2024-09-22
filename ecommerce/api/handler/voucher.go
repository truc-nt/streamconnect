package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IVoucherHandler interface {
	GetShopVouchers(ctx *gin.Context)
	AddVoucher(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type VoucherHandler struct {
	BaseHandler
	Service service.IVoucherService
}

func NewVoucherHandler(s service.IVoucherService) IVoucherHandler {
	return &VoucherHandler{
		Service: s,
	}
}

func (h *VoucherHandler) Create(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	var request *model.VoucherCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.Create(shopId, request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}

func (h *VoucherHandler) GetShopVouchers(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	vouchers, err := h.Service.GetShopVouchers(userId, shopId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, vouchers)
}

func (h *VoucherHandler) AddVoucher(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	voucherId, err := h.parseId(ctx, ctx.Param("voucher_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.AddVoucher(userId, voucherId); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}
