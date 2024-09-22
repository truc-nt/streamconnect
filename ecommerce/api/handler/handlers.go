package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler              IUserHandler
	ShopHandler              IShopHandler
	ProductHandler           IProductHandler
	VariantHandler           IVariantHandler
	ShopifyHandler           IShopifyHandler
	ExternalShopHandler      IExternalShopHandler
	ExternalVariantHandler   IExternalVariantHandler
	LivestreamHandler        ILivestreamHandler
	LivestreamProductHandler ILivestreamProductHandler
	CartHandler              ICartHandler
	OrderHandler             IOrderHandler
	VoucherHandler           IVoucherHandler
}

func ProvideHandlers(
	userHandler IUserHandler,
	shopHandler IShopHandler,
	productHandler IProductHandler,
	variantHandler IVariantHandler,
	shopifyHandler IShopifyHandler,
	externalShopHandler IExternalShopHandler,
	externalVariantHandler IExternalVariantHandler,
	livestreamHandler ILivestreamHandler,
	livestreamProductHandler ILivestreamProductHandler,
	cartHandler ICartHandler,
	orderHandler IOrderHandler,
	voucherHandler IVoucherHandler,
) *Handlers {
	return &Handlers{
		UserHandler:              userHandler,
		ShopHandler:              shopHandler,
		ProductHandler:           productHandler,
		VariantHandler:           variantHandler,
		ShopifyHandler:           shopifyHandler,
		ExternalShopHandler:      externalShopHandler,
		ExternalVariantHandler:   externalVariantHandler,
		LivestreamHandler:        livestreamHandler,
		LivestreamProductHandler: livestreamProductHandler,
		CartHandler:              cartHandler,
		OrderHandler:             orderHandler,
		VoucherHandler:           voucherHandler,
	}
}

type IBaseHandler interface {
	handleSuccessGet(ctx *gin.Context, data interface{})
	handleSuccessCreate(ctx *gin.Context)
	handleSuccessCreateWithData(ctx *gin.Context)
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

func (h *BaseHandler) handleSuccessCreateWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
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
