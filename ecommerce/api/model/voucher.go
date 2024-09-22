package model

import "time"

type VoucherCreateRequest struct {
	Code        string    `json:"code"`
	Discount    float64   `json:"discount"`
	MaxDiscount *float64  `json:"max_discount"`
	Method      string    `json:"method" xml:"method"`
	Type        string    `json:"type" xml:"type"`
	Target      string    `json:"target" xml:"target"`
	Quantity    int32     `json:"quantity" xml:"quantity"`
	MinPurchase *float64  `json:"min_purchase" xml:"min_spend"`
	StartTime   time.Time `json:"start_time" xml:"start_time"`
	EndTime     time.Time `json:"end_time" xml:"end_time"`
}
