package model

import "time"

type CreateLivestreamRequest struct {
	Title                      string                              `json:"title"`
	Description                string                              `json:"description"`
	StartTime                  string                              `json:"start_time"`
	EndTime                    *time.Time                          `json:"end_time"`
	LivestreamExternalVariants []*CreateLivestreamExternalVariants `json:"livestream_external_variants"`
}

type CreateLivestreamExternalVariants struct {
	IDProduct         int64 `json:"id_product"`
	IDVariant         int64 `json:"id_variant"`
	IDExternalVariant int64 `json:"id_external_variant"`
	Quantity          int32 `json:"quantity"`
}
