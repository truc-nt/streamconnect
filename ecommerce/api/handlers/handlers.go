package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ShopifyHandler IShopifyHandler
}

func NewHandlers(shopifyHandler IShopifyHandler) *Handlers {
	return &Handlers{
		ShopifyHandler: shopifyHandler,
	}
}

type IBaseHandler interface {
	HandleSuccessGet(ctx *gin.Context, data interface{})
	HandleSuccessCreate(ctx *gin.Context)
	HandleFailed(ctx *gin.Context, err error)
}

type BaseHandler struct{}

func (h *BaseHandler) HandleSuccessGet(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (h *BaseHandler) HandleSuccessCreate(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

func (h *BaseHandler) HandleSuccessUpdate(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (h *BaseHandler) HandleFailed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
