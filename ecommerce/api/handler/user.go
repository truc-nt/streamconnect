package handler

import (
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetDefaultAddress(ctx *gin.Context)
}

type UserHandler struct {
	BaseHandler

	Service service.IUserService
}

func NewUserHandler(s service.IUserService) IUserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) GetDefaultAddress(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.Param("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	address, err := h.Service.GetDefaultAddressByUserId(userId)

	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, address)
}
