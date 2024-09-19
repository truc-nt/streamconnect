package service

import (
	"database/sql"
	"ecommerce/api/model"
	"ecommerce/internal/adapter"
	"ecommerce/internal/constants"
	internalModel "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"errors"
	"strconv"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/samber/lo"
)

type ILivestreamService interface {
	CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error
	FetchLivestreams(status sql.NullString, ownerId sql.NullInt64) ([]internalModel.Livestream, error)
	GetLivestream(livestreamId int64) (*internalModel.Livestream, error)
	SetLivestreamHls(request *model.SetLivestreamHlsRequest) error
	RegisterLivestreamProductFollower(request *model.RegisterLivestreamProductFollowerRequest) error
}

type LivestreamService struct {
	LivestreamRepository                repository.ILivestreamRepository
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
	LivestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository
	VideoSdkAdapter                     adapter.IVideoSdkAdapter
}

func NewLivestreamService(
	livestreamService repository.ILivestreamRepository,
	livestreamProductRepository repository.ILivestreamProductRepository,
	livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository,
	videoSdkAdapter adapter.IVideoSdkAdapter,
	livestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository,
) ILivestreamService {
	return &LivestreamService{
		LivestreamRepository:                livestreamService,
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
		VideoSdkAdapter:                     videoSdkAdapter,
		LivestreamProductFollowerRepository: livestreamProductFollowerRepository,
	}
}

func (s *LivestreamService) CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) error {
	roomId, err := s.VideoSdkAdapter.CreateRoom()
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

func (s *LivestreamService) RegisterLivestreamProductFollower(request *model.RegisterLivestreamProductFollowerRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		//check if livestream product exists
		livestreamProducts, err := s.LivestreamProductRepository.FindAllLivestreamId(db, request.IDLivestream)
		if err != nil {
			return nil, err
		}
		livestreamProductIdsSet := make(map[int64]bool)
		for _, livestreamProduct := range livestreamProducts {
			livestreamProductIdsSet[livestreamProduct.IDLivestreamProduct] = true
		}
		newFollowers := make([]*internalModel.LivestreamProductFollower, len(request.IDLivestreamProducts))
		//create new followers
		for idx, livestreamProductId := range request.IDLivestreamProducts {
			if !livestreamProductIdsSet[livestreamProductId] {
				return nil, errors.New("livestream product with id " + strconv.FormatInt(livestreamProductId, 10) + " not found")
			}
			newLivestreamProductFollower := internalModel.LivestreamProductFollower{
				FkLivestreamProduct: livestreamProductId,
				FkUser:              request.IDUser,
				CreatedAt:           time.Now(),
			}
			newFollowers[idx] = &newLivestreamProductFollower
		}
		_, err = s.LivestreamProductFollowerRepository.CreateMany(
			db,
			postgres.ColumnList{
				table.LivestreamProductFollower.FkLivestreamProduct,
				table.LivestreamProductFollower.FkUser,
				table.LivestreamProductFollower.CreatedAt,
			},
			newFollowers,
		)
		return nil, nil
	}
	_, err := s.LivestreamProductFollowerRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}
	return nil
}
