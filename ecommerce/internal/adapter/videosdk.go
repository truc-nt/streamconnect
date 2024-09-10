package adapter

import (
	"ecommerce/internal/client/videosdk"
	"ecommerce/internal/configs"
)

type IVideoSdkAdapter interface {
	CreateRoom() (string, error)
}

type VideoSdkConfig struct {
	Token string
}

type VideoSdkAdapter struct {
	Config *VideoSdkConfig
}

func NewVideoSdkAdapter(config *configs.Config) IVideoSdkAdapter {
	return &VideoSdkAdapter{
		Config: &VideoSdkConfig{
			Token: config.VideoSdk.Token,
		},
	}
}

func (a *VideoSdkAdapter) getVideoSdkClient(param *videosdk.VideoSdkClientParam) videosdk.IVideoSdkClient {
	return videosdk.NewVideoSdkClient(param)
}

func (a *VideoSdkAdapter) CreateRoom() (string, error) {
	room, err := a.getVideoSdkClient(&videosdk.VideoSdkClientParam{
		Token: a.Config.Token,
	}).CreateRoom()

	if err != nil {
		return "", err
	}

	return room.RoomID, nil
}
