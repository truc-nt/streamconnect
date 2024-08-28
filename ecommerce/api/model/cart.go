package model

type AddToCartRequest struct {
	IDLivestreamExternalVariant int64 `json:"id_livestream_external_variant"`
	Quantity                    int32 `json:"quantity"`
}
