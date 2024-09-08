package service

import (
	"bytes"
	"database/sql"
	"ecommerce/api/model"
	"ecommerce/internal/constants"
	internalModel "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
	"io"
	"net/http"
)

type ILivestreamService interface {
	CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error
	FetchLivestreams(status sql.NullString, ownerId sql.NullInt64) ([]internalModel.Livestream, error)
	GetLivestream(livestreamId int64) (*internalModel.Livestream, error)
	SetLivestreamHls(request *model.SetLivestreamHlsRequest) error
}

type LivestreamService struct {
	LivestreamRepository                repository.ILivestreamRepository
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
}

func NewLivestreamService(livestreamService repository.ILivestreamRepository, livestreamProductRepository repository.ILivestreamProductRepository, livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository) ILivestreamService {
	return &LivestreamService{
		LivestreamRepository:                livestreamService,
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
	}
}

const videoSdkBaseUrl = "https://api.videosdk.live/v2"
const videoSdkToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcGlrZXkiOiJjN2MwOTgwMy05OWUzLTRmMGUtOTg3Ny0zYjU1MTdiNThkY2IiLCJwZXJtaXNzaW9ucyI6WyJhbGxvd19qb2luIl0sImlhdCI6MTcyMzMyOTgyOSwiZXhwIjoxNzI1OTIxODI5fQ.5x11sT5M7jzIM9EslqanSiMpnLeLTImr-zlzDKUuntc"

func createVideoSdkRoom() (string, error) {
	client := &http.Client{}

	reqBody := []byte(`{}`)
	req, err := http.NewRequest("POST", videoSdkBaseUrl+"/rooms", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", videoSdkToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response body: %w", err)
	}

	roomId, ok := result["roomId"].(string)
	if !ok {
		return "", fmt.Errorf("roomId not found in response")
	}

	return roomId, nil
}

func (s *LivestreamService) CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error {
	roomId, err := createVideoSdkRoom()
	if err != nil {
		return err
	}
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {

		newLivestreamData := internalModel.Livestream{
			FkShop:      shopId,
			Title:       createLivestreamRequest.Title,
			Description: &createLivestreamRequest.Description,
			Status:      constants.LIVESTREAM_CREATED,
			MeetingID:   roomId,
		}
		newLivestream, err := s.LivestreamRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Livestream.FkShop,
				table.Livestream.Title,
				table.Livestream.Description,
				table.Livestream.Status,
				table.Livestream.MeetingID,
			},
			newLivestreamData,
		)
		if err != nil {
			return nil, err
		}

		for _, livestreamProduct := range createLivestreamRequest.LivestreamProducts {
			newLivestreamProductData := internalModel.LivestreamProduct{
				FkLivestream: newLivestream.IDLivestream,
				FkProduct:    livestreamProduct.IDProduct,
				Priority:     livestreamProduct.Priority,
			}
			newLivestreamProduct, err := s.LivestreamProductRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.LivestreamProduct.FkLivestream,
					table.LivestreamProduct.FkProduct,
					table.LivestreamProduct.Priority,
				},
				newLivestreamProductData,
			)
			if err != nil {
				return nil, err
			}

			newExternalLivestreamVariantData := make([]*internalModel.LivestreamExternalVariant, 0)
			for _, livestreamVariant := range livestreamProduct.LivestreamVariants {
				livestreamExternalVariants := lo.Map(livestreamVariant.LivestreamExternalVariants, func(externalVariant *struct {
					IDExternalVariant int64 `json:"id_external_variant"`
					Quantity          int32 `json:"quantity"`
				}, index int) *internalModel.LivestreamExternalVariant {
					return &internalModel.LivestreamExternalVariant{
						FkLivestreamProduct: newLivestreamProduct.IDLivestreamProduct,
						FkExternalVariant:   externalVariant.IDExternalVariant,
						Quantity:            externalVariant.Quantity,
					}
				})
				newExternalLivestreamVariantData = append(newExternalLivestreamVariantData, livestreamExternalVariants...)
			}

			_, err = s.LivestreamExternalVariantRepository.CreateMany(
				db,
				postgres.ColumnList{
					table.LivestreamExternalVariant.FkLivestreamProduct,
					table.LivestreamExternalVariant.FkExternalVariant,
					table.LivestreamExternalVariant.Quantity,
				},
				newExternalLivestreamVariantData,
			)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err = s.LivestreamRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *LivestreamService) FetchLivestreams(status sql.NullString, ownerId sql.NullInt64) ([]internalModel.Livestream, error) {
	livestreams, err := s.LivestreamRepository.GetByStatusAndOwnerId(s.LivestreamRepository.GetDatabase().Db, status, ownerId)
	if err != nil {
		return nil, err
	}
	return livestreams, nil
}

func (s *LivestreamService) GetLivestream(livestreamId int64) (*internalModel.Livestream, error) {
	livestream, err := s.LivestreamRepository.GetById(s.LivestreamRepository.GetDatabase().Db, livestreamId)
	if err != nil {
		return nil, err
	}
	return livestream, nil
}

func (s *LivestreamService) SetLivestreamHls(request *model.SetLivestreamHlsRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		livestream, err := s.LivestreamRepository.GetById(db, request.IDLivestream)
		if err != nil {
			return nil, err
		}

		livestream.HlsURL = &request.HlsUrl
		livestream.Status = constants.LIVESTREAM_STREAMING
		_, err = s.LivestreamRepository.UpdateById(
			db,
			postgres.ColumnList{
				table.Livestream.HlsURL,
				table.Livestream.Status,
			},
			*livestream,
		)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := s.LivestreamRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}
