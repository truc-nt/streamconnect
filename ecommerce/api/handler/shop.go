package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type IShopHandler interface {
	CreateShopForNewAccount(ctx *gin.Context)
}

type ShopHandler struct {
	BaseHandler
	Service service.IShopService
}

func NewShopHandler(s service.IShopService) IShopHandler {
	return &ShopHandler{
		Service: s,
	}
}

func (h *ShopHandler) CreateShopForNewAccount(ctx *gin.Context) {
	//userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	//if err != nil {
	//	h.handleFailed(ctx, errors.New("error while parsing user_id: " + err.Error()))
	//	return
	//}
	var request *model.CreateShopForNewUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	//if userId != request.UserID {
	//	//later can expand to one request to create shop for other user
	//	h.handleFailed(ctx, errors.New("user_id in header and body are different")
	//	return
	//}
	if err := h.Service.CreateShopForNewAccount(request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessCreate(ctx)
}
