package model

import "github.com/jackc/pgtype"

type GetVariantsByProductIdParam struct {
	PaginationParam
}

type GetVariantsByProductIdData struct {
	Name     string      `json:"name"`
	Status   string      `json:"status"`
	Option   pgtype.JSON `json:"option"`
	Variants []*struct {
		IDVariant int64       `json:"id_variant"`
		Sku       string      `json:"sku"`
		Status    string      `json:"status"`
		Option    pgtype.JSON `json:"option"`
		Stock     int64       `json:"stock"`
		Price     float64     `json:"price"`
	}
}

type ConnectVariantsRequest []*struct {
	IDVariant         int64 `json:"id_variant"`
	IDExternalVariant int64 `json:"id_external_variant"`
}
