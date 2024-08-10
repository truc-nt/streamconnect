package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-viper/mapstructure/v2"
)

type IExternalProductHandler interface {
	GetExternalProductsByExternalShopId(ctx *gin.Context)
}

type ExternalProductHandler struct {
	BaseHandler
	Service service.IExternalProductService
}

func NewExternalProductHandler(s service.IExternalProductService) IExternalProductHandler {
	return &ExternalProductHandler{
		Service: s,
	}
}

func (h *ExternalProductHandler) GetExternalProductsByExternalShopId(ctx *gin.Context) {
	externalShopId, err := h.parseId(ctx, ctx.Param("external_shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var queryParams *model.GetExternalProductsByExternalShopIdParam
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	products, err := h.Service.GetExternalProductsByExternalShopId(externalShopId, queryParams.Limit, queryParams.Offset)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var res []*model.GetExternalProductsByExternalShopIdData
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &res,
	})
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := decoder.Decode(products); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	h.handleSuccessGet(ctx, res)
}
