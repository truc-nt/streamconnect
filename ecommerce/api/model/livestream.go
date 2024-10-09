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

type UpdateLivestreamRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
	MeetingID   *string    `json:"meeting_id"`
	HlsURL      *string    `json:"hls_url"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *string    `json:"end_time"`
}

type SetLivestreamHlsRequest struct {
	HlsUrl string `json:"hls_url"`
}

type GetLivestreamResponse struct {
	IDLivestream int64   `json:"id_livestream" xml:"id_livestream"`
	IDShop       int64   `json:"id_shop" xml:"fk_shop"`
	Title        string  `json:"title" xml:"title"`
	Description  *string `json:"description" xml:"description"`
	Status       string  `json:"status" xml:"status"`
	MeetingID    string  `json:"meeting_id" xml:"meeting_id"`
	HlsURL       *string `json:"hls_url" xml:"hls_url"`
	IsHost       bool    `json:"is_host"`
	ShopName     string  `json:"shop_name"`
	IsFollowed   bool    `json:"is_followed"`
}

type UpdateLivestreamExternalVariantQuantityRequest []*struct {
	LivestreamExternalVariantId int64 `json:"id_livestream_external_variant"`
	Quantity                    int32 `json:"quantity"`
}

type UpdateLivestreamProductPriorityRequest []*struct {
	LivestreamProductId int64 `json:"id_livestream_product"`
	Priority            int32 `json:"priority"`
}

type UpdateLivestreamProductsRequest []*struct {
	IDLivestreamProduct int64  `json:"id_livestream_product"`
	Priority            *int32 `json:"priority"`
	IsLivestreamed      *bool  `json:"is_livestreamed"`
}

type RegisterLivestreamProductFollowerRequest struct {
	IDLivestreamProducts []int64 `json:"id_livestream_products"`
	IDUser               int64   `json:"id_user"`
}

type LivestreamProductFollowerDTO struct {
	UserIds    []int64           `json:"user_ids"`
	Livestream *model.Livestream `json:"livestream"`
	Product    *model.Product    `json:"livestream_product"`
}
