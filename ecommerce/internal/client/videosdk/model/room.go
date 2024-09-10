package model

import "time"

type CreateRoomResponse struct {
	RoomID       string    `json:"roomId"`
	CustomRoomID string    `json:"customRoomId"`
	UserID       string    `json:"userId"`
	Disabled     bool      `json:"disabled"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ID           string    `json:"id"`
	Links        struct {
		GetRoom    string `json:"get_room"`
		GetSession string `json:"get_session"`
	} `json:"links"`
}
