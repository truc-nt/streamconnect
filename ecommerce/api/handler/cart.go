package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type ICartHandler interface {
	Get(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
}

type CartHandler struct {
	BaseHandler
	Service service.ICartService
}

func NewCartHandler(s service.ICartService) ICartHandler {
	return &CartHandler{
		Service: s,
	}
}

func (h *CartHandler) Get(ctx *gin.Context) {
	cartId, err := h.parseId(ctx, ctx.Param("cart_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	cart, err := h.Service.Get(cartId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, cart)
}

func (h *CartHandler) AddToCart(ctx *gin.Context) {
	cartId, err := h.parseId(ctx, ctx.Param("cart_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var addToCart *model.AddToCartRequest
	if err := ctx.ShouldBindJSON(&addToCart); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.AddToCart(cartId, addToCart); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessUpdate(ctx)
}
