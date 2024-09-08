package videosdk

import (
	"ecommerce/internal/client/videosdk/model"
	"ecommerce/internal/constants"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IVideoSdkClient interface {
	CreateRoom() (*model.CreateRoomResponse, error)
}

type VideoSdkClientParam struct {
	Token string
}

type VideoSdkClient struct {
	Client *http.Client
	Param  *VideoSdkClientParam
}

func NewVideoSdkClient(param *VideoSdkClientParam) IVideoSdkClient {
	return &VideoSdkClient{
		Client: &http.Client{},
		Param:  param,
	}
}

func (c *VideoSdkClient) getResponse(req *http.Request) ([]byte, error) {
	req.Header.Set(constants.VideoSdkTokenKey, c.Param.Token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func (c *VideoSdkClient) CreateRoom() (*model.CreateRoomResponse, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", constants.VideoSdkBaseUrl, constants.VideoSdkCreateRoomPath), nil)
	if err != nil {
		return nil, err
	}

	resData, err := c.getResponse(req)
	if err != nil {
		return nil, err
	}

	var response *model.CreateRoomResponse
	if err := json.Unmarshal(resData, &response); err != nil {
		return nil, err
	}
	return response, nil
}
