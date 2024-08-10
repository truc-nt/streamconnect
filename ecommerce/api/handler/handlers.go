package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ProductHandler         IProductHandler
	VariantHandler         IVariantHandler
	ShopifyHandler         IShopifyHandler
	ExternalShopHandler    IExternalShopHandler
	ExternalProductHandler IExternalProductHandler
}

func ProvideHandlers(
	productHandler IProductHandler,
	variantHandler IVariantHandler,
	shopifyHandler IShopifyHandler,
	externalShopHandler IExternalShopHandler,
	externalProductHandler IExternalProductHandler,
) *Handlers {
	return &Handlers{
		ProductHandler:         productHandler,
		VariantHandler:         variantHandler,
		ShopifyHandler:         shopifyHandler,
		ExternalShopHandler:    externalShopHandler,
		ExternalProductHandler: externalProductHandler,
	}
}

type IBaseHandler interface {
	handleSuccessGet(ctx *gin.Context, data interface{})
	handleSuccessCreate(ctx *gin.Context)
	handleSuccessUpdate(ctx *gin.Context)
	handleFailed(ctx *gin.Context, err error)
	parseId(c *gin.Context, id string) int64
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

func (h *BaseHandler) parseId(ctx *gin.Context, id string) (int64, error) {
	_id, err := strconv.ParseInt(id, 10, 64)
	return _id, err
}
