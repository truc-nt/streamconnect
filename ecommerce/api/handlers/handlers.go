package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ShopifyHandler      IShopifyHandler
	ExternalShopHandler IExternalShopHandler
}

func NewHandlers(shopifyHandler IShopifyHandler, externalShopHandler IExternalShopHandler) *Handlers {
	return &Handlers{
		ShopifyHandler:      shopifyHandler,
		ExternalShopHandler: externalShopHandler,
	}
}

type IBaseHandler interface {
	handleSuccessGet(ctx *gin.Context, data interface{})
	handleSuccessCreate(ctx *gin.Context)
	handleSuccessUpdate(ctx *gin.Context)
	handleFailed(ctx *gin.Context, err error)
	parseId(c *gin.Context, id string) int32
}

type BaseHandler struct{}

func (h *BaseHandler) handleSuccessGet(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (h *BaseHandler) handleSuccessCreate(ctx *gin.Context) {
	ctx.Status(http.StatusCreated)
}

func (h *BaseHandler) handleSuccessUpdate(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (h *BaseHandler) handleFailed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *BaseHandler) parseId(ctx *gin.Context, id string) (int32, error) {
	_id, err := strconv.ParseInt(id, 10, 32)
	return int32(_id), err
}
