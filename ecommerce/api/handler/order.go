package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IOrderHandler interface {
	CreateOrderWithCartItems(ctx *gin.Context)
}

type OrderHandler struct {
	BaseHandler
	Service service.IOrderService
}

func NewOrderHandler(s service.IOrderService) IOrderHandler {
	return &OrderHandler{
		Service: s,
	}
}

func (h *OrderHandler) CreateOrderWithCartItems(ctx *gin.Context) {
	var createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest
	if err := ctx.ShouldBindJSON(&createOrderWithCartItemsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	err := h.Service.CreateOrderWithCartItems(createOrderWithCartItemsRequest)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
}
