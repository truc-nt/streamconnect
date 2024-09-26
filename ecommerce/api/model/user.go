package model

import "time"

type UpdateUserRequest struct {
	Email     *string    `json:"email"`
	Gender    *string    `json:"gender"`
	Birthdate *time.Time `json:"birthdate"`
}

type AddressCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
}
