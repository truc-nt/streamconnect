package model

import (
	"ecommerce/internal/database/gen/model"
	"time"
)

type GetLivestreamsQueryParam struct {
	Status []string `form:"status[]"`
	ShopId int64    `form:"shop_id,default=0"`
	PaginationParam
}

type LivestreamProductCreateRequest struct {
	IDProduct          int64 `json:"id_product"`
	Priority           int32 `json:"priority"`
	LivestreamVariants []*struct {
		IDVariant                  int64 `json:"id_variant"`
		LivestreamExternalVariants []*struct {
			IDExternalVariant int64 `json:"id_external_variant"`
			Quantity          int32 `json:"quantity"`
		} `json:"livestream_external_variants"`
	} `json:"livestream_variants"`
}

type CreateLivestreamRequest struct {
	Title              string                            `json:"title"`
	Description        string                            `json:"description"`
	StartTime          *time.Time                        `json:"start_time"`
	EndTime            *time.Time                        `json:"end_time"`
	LivestreamProducts []*LivestreamProductCreateRequest `json:"livestream_products"`
}

type SetLivestreamHlsRequest struct {
	HlsUrl string `json:"hls_url"`
}

type GetLivestreamInfoResponse struct {
	MeetingId string `json:"meeting_id"`
	IDShop    int64  `json:"id_shop"`
	IsHost    bool   `json:"is_host"`
	ShopName  string `json:"shop_name"`
}

type UpdateLivestreamExternalVariantQuantityRequest []*struct {
	LivestreamExternalVariantId int64 `json:"id_livestream_external_variant"`
	Quantity                    int32 `json:"quantity"`
}

type PinLivestreamProductRequest []struct {
	LivestreamProductId int64 `json:"id_livestream_product"`
	Priority            int32 `json:"priority"`
}

type RegisterLivestreamProductFollowerRequest struct {
	IDLivestreamProducts []int64 `json:"id_livestream_products"`
	IDLivestream         int64   `json:"id_livestream"`
	IDUser               int64   `json:"id_user"`
}

type LivestreamProductFollowerDTO struct {
	UserIds    []int64           `json:"user_ids"`
	Livestream *model.Livestream `json:"livestream"`
	Product    *model.Product    `json:"livestream_product"`
}
