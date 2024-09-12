package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetDefaultAddress(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetAddressesByUserId(ctx *gin.Context)
	CreateAddress(ctx *gin.Context)
}

type UserHandler struct {
	BaseHandler

	UserService        service.IUserService
	UserAddressService service.IUserAddressService
}

func NewUserHandler(userService service.IUserService, userAddressService service.IUserAddressService) IUserHandler {
	return &UserHandler{
		UserService:        userService,
		UserAddressService: userAddressService,
	}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	user, err := h.UserService.GetByUserId(userId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, user)
}

func (h *UserHandler) GetDefaultAddress(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	address, err := h.UserService.GetDefaultAddressByUserId(userId)

	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, address)
}

func (h *UserHandler) GetAddressesByUserId(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	addresses, err := h.UserAddressService.GetAddressByUserId(userId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, addresses)
}

func (h *UserHandler) CreateAddress(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var address *model.AddressCreateRequest
	if err := ctx.ShouldBindJSON(&address); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if _, err := h.UserAddressService.CreateAddress(userId, address); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessCreate(ctx)
}
