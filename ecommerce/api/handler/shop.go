package handler

import (
	"ecommerce/api/model"
	"ecommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type IShopHandler interface {
	CreateShopForNewAccount(ctx *gin.Context)
	GetShop(ctx *gin.Context)
	UpdateShop(ctx *gin.Context)
	FollowShop(ctx *gin.Context)
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

func (h *ShopHandler) GetShop(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
	}

	isFollowing := false

	if ctx.GetHeader("user_id") != "" {
		userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
		if err != nil {
			h.handleFailed(ctx, err)
			return
		}

		isFollowing, err = h.Service.IsFollowed(shopId, userId)
		if err != nil {
			h.handleFailed(ctx, err)
			return
		}
	}

	shop, err := h.Service.GetShop(shopId)
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessGet(ctx, model.GetShopResponse{
		IDShop:      shop.IDShop,
		FkUser:      shop.FkUser,
		Name:        shop.Name,
		Description: shop.Description,
		CreatedAt:   shop.CreatedAt,
		UpdatedAt:   shop.UpdatedAt,
		IsFollowing: isFollowing,
	})
}

func (h *ShopHandler) UpdateShop(ctx *gin.Context) {
	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	var request *model.UpdateShopRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.UpdateShop(shopId, request); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}

func (h *ShopHandler) FollowShop(ctx *gin.Context) {
	userId, err := h.parseId(ctx, ctx.GetHeader("user_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	shopId, err := h.parseId(ctx, ctx.Param("shop_id"))
	if err != nil {
		h.handleFailed(ctx, err)
		return
	}

	if err := h.Service.AddFollower(shopId, userId); err != nil {
		h.handleFailed(ctx, err)
		return
	}
	h.handleSuccessUpdate(ctx)
}
