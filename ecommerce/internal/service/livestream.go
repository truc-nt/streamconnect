package service

import (
	"ecommerce/api/model"
	"ecommerce/internal/adapter"
	"ecommerce/internal/constants"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	"github.com/samber/lo"
)

type ILivestreamService interface {
	CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) (int64, error)
	GetLivestreams(param *model.GetLivestreamsQueryParam) ([]*entity.Livestream, error)
	GetLivestream(livestreamId int64) (*repository.GetInfo, error)
	SetLivestreamHls(livestreamId int64, hlsUrl string) error
	UpdateLivestream(livestreamId int64, request *model.UpdateLivestreamRequest) error
	UpdateLivestreamExternalVariantQuantity(updateLivestreamExternalVariantQuantityRequest *model.UpdateLivestreamExternalVariantQuantityRequest) error
	AddLivestreamProduct(livestreamId int64, livestreamProductCreateRequest []*model.LivestreamProductCreateRequest) error
	StartLivestream(livestreamId int64) error
	RegisterLivestreamProductFollower(livestreamId int64, userId int64, livestreamProductIds []int64) error
	FetchLivestreamProductFollowers(productId int64) (*model.LivestreamProductFollowerDTO, error)
	UpdateLivestreamProducts(updateLivestreamProductsRequest *model.UpdateLivestreamProductsRequest) error
}

type LivestreamService struct {
	LivestreamRepository                repository.ILivestreamRepository
	LivestreamProductRepository         repository.ILivestreamProductRepository
	LivestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository
	LivestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository
	ProductRepository                   repository.IProductRepository
	VideoSdkAdapter                     adapter.IVideoSdkAdapter
}

func NewLivestreamService(
	livestreamService repository.ILivestreamRepository,
	livestreamProductRepository repository.ILivestreamProductRepository,
	livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository,
	videoSdkAdapter adapter.IVideoSdkAdapter,
	livestreamProductFollowerRepository repository.ILivestreamProductFollowerRepository,
	productRepository repository.IProductRepository,
) ILivestreamService {
	return &LivestreamService{
		LivestreamRepository:                livestreamService,
		LivestreamProductRepository:         livestreamProductRepository,
		LivestreamExternalVariantRepository: livestreamExternalVariantRepository,
		VideoSdkAdapter:                     videoSdkAdapter,
		LivestreamProductFollowerRepository: livestreamProductFollowerRepository,
		ProductRepository:                   productRepository,
	}
}

func (s *LivestreamService) CreateLivestream(shopId int64, createLivestreamRequest *model.CreateLivestreamRequest) (int64, error) {
	roomId, err := s.VideoSdkAdapter.CreateRoom()
	if err != nil {
		return 0, err
	}
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		status := constants.LIVESTREAM_CREATED
		if createLivestreamRequest.StartTime == nil {
			status = constants.LIVESTREAM_STARTED
		}

		newLivestreamData := entity.Livestream{
			FkShop:      shopId,
			Title:       createLivestreamRequest.Title,
			Description: &createLivestreamRequest.Description,
			Status:      status,
			MeetingID:   roomId,
		}

		columns := postgres.ColumnList{
			table.Livestream.FkShop,
			table.Livestream.Title,
			table.Livestream.Description,
			table.Livestream.Status,
			table.Livestream.MeetingID,
		}

		if createLivestreamRequest.StartTime != nil {
			newLivestreamData.StartTime = *createLivestreamRequest.StartTime
			columns = append(columns, table.Livestream.StartTime)
		}
		newLivestream, err := s.LivestreamRepository.CreateOne(
			db,
			columns,
			newLivestreamData,
		)
		if err != nil {
			return nil, err
		}

		for _, livestreamProduct := range createLivestreamRequest.LivestreamProducts {
			newLivestreamProductData := entity.LivestreamProduct{
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

			newExternalLivestreamVariantData := make([]*entity.LivestreamExtVariant, 0)
			for _, livestreamVariant := range livestreamProduct.LivestreamVariants {
				livestreamExternalVariants := lo.Map(livestreamVariant.LivestreamExternalVariants, func(externalVariant *struct {
					IDExternalVariant int64 `json:"id_external_variant"`
					Quantity          int32 `json:"quantity"`
				}, index int) *entity.LivestreamExtVariant {
					return &entity.LivestreamExtVariant{
						FkLivestreamProduct: newLivestreamProduct.IDLivestreamProduct,
						FkExtVariant:        externalVariant.IDExternalVariant,
						Quantity:            externalVariant.Quantity,
					}
				})
				newExternalLivestreamVariantData = append(newExternalLivestreamVariantData, livestreamExternalVariants...)
			}

			_, err = s.LivestreamExternalVariantRepository.CreateMany(
				db,
				postgres.ColumnList{
					table.LivestreamExtVariant.FkLivestreamProduct,
					table.LivestreamExtVariant.FkExtVariant,
					table.LivestreamExtVariant.Quantity,
				},
				newExternalLivestreamVariantData,
			)
			if err != nil {
				return nil, err
			}
		}

		return newLivestream.IDLivestream, nil
	}

	result, err := s.LivestreamRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return 0, err
	}

	livestreamId, ok := result.(int64)
	if !ok {
		return 0, fmt.Errorf("unexpected type for livestream ID")
	}

	return livestreamId, nil
}

func (s *LivestreamService) GetLivestreams(param *model.GetLivestreamsQueryParam) ([]*entity.Livestream, error) {
	livestreams, err := s.LivestreamRepository.GetByParam(s.LivestreamRepository.GetDatabase().Db, param)
	if err != nil {
		return nil, err
	}
	return livestreams, nil
}

func (s *LivestreamService) GetLivestream(livestreamId int64) (*repository.GetInfo, error) {
	livestream, err := s.LivestreamRepository.GetInfoById(s.LivestreamRepository.GetDatabase().Db, livestreamId)
	if err != nil {
		return nil, err
	}
	return livestream, nil
}

func (s *LivestreamService) UpdateLivestream(livestreamId int64, request *model.UpdateLivestreamRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		livestream, err := s.LivestreamRepository.GetById(db, livestreamId)
		if err != nil {
			return nil, err
		}

		columnList := postgres.ColumnList{}
		if request.Title != nil {
			livestream.Title = *request.Title
			columnList = append(columnList, table.Livestream.Title)
		}
		if request.Description != nil {
			livestream.Description = request.Description
			columnList = append(columnList, table.Livestream.Description)
		}
		if request.StartTime != nil {
			livestream.StartTime = *request.StartTime
			columnList = append(columnList, table.Livestream.StartTime)
		}
		if request.Status != nil {
			livestream.Status = *request.Status
			columnList = append(columnList, table.Livestream.Status)
		}

		_, err = s.LivestreamRepository.UpdateById(
			db,
			columnList,
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

func (s *LivestreamService) SetLivestreamHls(livestreamId int64, hlsUrl string) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		livestream, err := s.LivestreamRepository.GetById(db, livestreamId)
		if err != nil {
			return nil, err
		}

		livestream.HlsURL = &hlsUrl
		livestream.Status = constants.LIVESTREAM_PLAYED
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

func (s *LivestreamService) RegisterLivestreamProductFollower(livestreamId int64, userId int64, livestreamProductIds []int64) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		livestreamProducts, err := s.LivestreamProductRepository.FindAllLivestreamId(db, livestreamId)
		if err != nil {
			return nil, err
		}
		livestreamProductIdsSet := make(map[int64]bool)
		for _, livestreamProduct := range livestreamProducts {
			livestreamProductIdsSet[livestreamProduct.IDLivestreamProduct] = true
		}
		newFollowers := make([]*entity.LivestreamProductFollower, len(livestreamProductIds))
		//create new followers
		for idx, livestreamProductId := range livestreamProductIds {
			if !livestreamProductIdsSet[livestreamProductId] {
				return nil, errors.New("livestream product with id " + strconv.FormatInt(livestreamProductId, 10) + " not found")
			}
			newLivestreamProductFollower := entity.LivestreamProductFollower{
				FkLivestreamProduct: livestreamProductId,
				FkUser:              userId,
				CreatedAt:           time.Now(),
			}
			newFollowers[idx] = &newLivestreamProductFollower
		}
		if _, err := s.LivestreamProductFollowerRepository.CreateMany(
			db,
			postgres.ColumnList{
				table.LivestreamProductFollower.FkLivestreamProduct,
				table.LivestreamProductFollower.FkUser,
				table.LivestreamProductFollower.CreatedAt,
			},
			newFollowers,
		); err != nil {
			return nil, err
		}

		return nil, nil
	}
	_, err := s.LivestreamProductFollowerRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *LivestreamService) UpdateLivestreamExternalVariantQuantity(updateLivestreamExternalVariantQuantityRequest *model.UpdateLivestreamExternalVariantQuantityRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, product := range *updateLivestreamExternalVariantQuantityRequest {
			if _, err := s.LivestreamExternalVariantRepository.UpdateById(
				db,
				postgres.ColumnList{
					table.LivestreamExtVariant.Quantity,
				},
				entity.LivestreamExtVariant{
					IDLivestreamExtVariant: product.LivestreamExternalVariantId,
					Quantity:               product.Quantity,
				},
			); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.LivestreamExternalVariantRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *LivestreamService) AddLivestreamProduct(livestreamId int64, livestreamProductCreateRequest []*model.LivestreamProductCreateRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, request := range livestreamProductCreateRequest {
			livestreamProduct, err := s.LivestreamProductRepository.GetByLivestreamIdAndProductId(db, livestreamId, request.IDProduct)
			if err != nil {
				var pgErr pgx.PgError
				if errors.As(err, &pgErr) {
					if pgErr.Code != pgerrcode.NoDataFound {
						return nil, err
					}
				}

				livestreamProduct, err = s.LivestreamProductRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.LivestreamProduct.FkLivestream,
						table.LivestreamProduct.FkProduct,
						table.LivestreamProduct.Priority,
					},
					entity.LivestreamProduct{
						FkLivestream: livestreamId,
						FkProduct:    request.IDProduct,
						Priority:     request.Priority,
					},
				)

				if err != nil {
					return nil, err
				}
			}

			isLivestreamed := false
			if _, err := s.LivestreamProductRepository.UpdateById(
				db,
				postgres.ColumnList{
					table.LivestreamProduct.IsLivestreamed,
				},
				entity.LivestreamProduct{
					IDLivestreamProduct: livestreamProduct.IDLivestreamProduct,
					IsLivestreamed:      &isLivestreamed,
				},
			); err != nil {
				return nil, err
			}

			newExternalLivestreamVariantData := make([]*entity.LivestreamExtVariant, 0)
			for _, livestreamVariant := range request.LivestreamVariants {
				livestreamExternalVariants := lo.Map(livestreamVariant.LivestreamExternalVariants, func(externalVariant *struct {
					IDExternalVariant int64 `json:"id_external_variant"`
					Quantity          int32 `json:"quantity"`
				}, index int) *entity.LivestreamExtVariant {
					return &entity.LivestreamExtVariant{
						FkLivestreamProduct: livestreamProduct.IDLivestreamProduct,
						FkExtVariant:        externalVariant.IDExternalVariant,
						Quantity:            externalVariant.Quantity,
					}
				})
				newExternalLivestreamVariantData = append(newExternalLivestreamVariantData, livestreamExternalVariants...)
			}

			_, err = s.LivestreamExternalVariantRepository.CreateManyOnConflict(
				db,
				postgres.ColumnList{
					table.LivestreamExtVariant.FkLivestreamProduct,
					table.LivestreamExtVariant.FkExtVariant,
					table.LivestreamExtVariant.Quantity,
				},
				newExternalLivestreamVariantData,
			)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.LivestreamExternalVariantRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *LivestreamService) StartLivestream(livestreamId int64) error {
	if _, err := s.LivestreamRepository.UpdateById(
		s.LivestreamRepository.GetDatabase().Db,
		postgres.ColumnList{
			table.Livestream.Status,
		},
		entity.Livestream{
			IDLivestream: livestreamId,
			Status:       constants.LIVESTREAM_STARTED,
		},
	); err != nil {
		return err
	}

	return nil
}

func (s *LivestreamService) FetchLivestreamProductFollowers(productId int64) (*model.LivestreamProductFollowerDTO, error) {
	followers, err := s.LivestreamProductFollowerRepository.FindByProductId(
		s.LivestreamProductFollowerRepository.GetDatabase().Db,
		productId,
	)
	if err != nil {
		return nil, err
	}
	if len(followers) == 0 {
		return &model.LivestreamProductFollowerDTO{}, nil
	}
	//extract user ids
	var userIds = make([]int64, len(followers))
	for idx, follower := range followers {
		userIds[idx] = follower.FkUser
	}
	livestreamProduct, err := s.LivestreamProductRepository.GetById(s.LivestreamProductRepository.GetDatabase().Db, productId)
	if err != nil {
		return nil, err
	}
	//fetch livestream
	livestream, err := s.LivestreamRepository.GetById(s.LivestreamRepository.GetDatabase().Db, livestreamProduct.FkLivestream)
	if err != nil {
		return nil, err
	}
	//fetch product
	product, err := s.ProductRepository.GetById(s.ProductRepository.GetDatabase().Db, livestreamProduct.FkProduct)
	if err != nil {
		return nil, err
	}
	var data = model.LivestreamProductFollowerDTO{}
	data.UserIds = userIds
	data.Product = product
	data.Livestream = livestream

	return &data, nil
}

func (s *LivestreamService) UpdateLivestreamProducts(updateLivestreamProductsRequest *model.UpdateLivestreamProductsRequest) error {
	var execWithinTransaction = func(db qrm.Queryable) (interface{}, error) {
		for _, request := range *updateLivestreamProductsRequest {
			updatedColumnList := postgres.ColumnList{}
			livestreamProduct := entity.LivestreamProduct{
				IDLivestreamProduct: request.IDLivestreamProduct,
			}
			if request.IsLivestreamed != nil {
				updatedColumnList = append(
					updatedColumnList,
					table.LivestreamProduct.IsLivestreamed,
				)
				livestreamProduct.IsLivestreamed = request.IsLivestreamed
			}
			if request.Priority != nil {
				updatedColumnList = append(
					updatedColumnList,
					table.LivestreamProduct.Priority,
				)
				livestreamProduct.Priority = *request.Priority
			}

			if _, err := s.LivestreamProductRepository.UpdateById(
				db,
				updatedColumnList,
				livestreamProduct,
			); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	_, err := s.LivestreamProductRepository.ExecWithinTransaction(execWithinTransaction)
	if err != nil {
		return err
	}

	return nil
}
