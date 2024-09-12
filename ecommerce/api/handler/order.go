package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IOrderHandler interface {
	CreateOrderWithCartItems(ctx *gin.Context)
	GetBuyOrders(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
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
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	var createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest
	if err := ctx.ShouldBindJSON(&createOrderWithCartItemsRequest); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	err = h.Service.CreateOrderWithCartItems(userId, createOrderWithCartItemsRequest)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
}

func (h *OrderHandler) GetBuyOrders(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	orders, err := h.Service.GetBuyOrders(userId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, orders)
}

func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	orderId, err := h.parseId(ctx, ctx.Param("order_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	order, err := h.Service.GetOrder(orderId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, order)
}
