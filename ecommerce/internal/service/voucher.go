package service

import (
	apiModel "ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	"ecommerce/internal/repository"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
)

type IVoucherService interface {
	Create(shopId int64, request *apiModel.VoucherCreateRequest) error
	GetShopVouchers(userId, shopId int64) (interface{}, error)
	AddVoucher(userId int64, voucherId int64) error
}

type VoucherService struct {
	VoucherRepository     repository.IVoucherRepository
	VoucherUserRepository repository.IVoucherUserRepository
}

func NewVoucherService(
	voucherRepository repository.IVoucherRepository,
	voucherUserRepository repository.IVoucherUserRepository,
) IVoucherService {
	return &VoucherService{
		VoucherRepository:     voucherRepository,
		VoucherUserRepository: voucherUserRepository,
	}
}

func (s *VoucherService) Create(shopId int64, request *apiModel.VoucherCreateRequest) error {
	newData := entity.Voucher{
		FkShop:      shopId,
		Code:        request.Code,
		Discount:    request.Discount,
		MaxDiscount: request.MaxDiscount,
		Method:      request.Method,
		Type:        request.Type,
		Target:      request.Target,
		Quantity:    request.Quantity,
		MinPurchase: *request.MinPurchase,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
	}
	if _, err := s.VoucherRepository.CreateOne(
		s.VoucherRepository.GetDatabase().Db,
		table.Voucher.MutableColumns,
		newData,
	); err != nil {
		return err
	}

	return nil
}

func (s *VoucherService) GetShopVouchers(userId, shopId int64) (interface{}, error) {
	return s.VoucherRepository.GetShopVouchers(s.VoucherRepository.GetDatabase().Db, userId, shopId)
}

func (s *VoucherService) AddVoucher(userId int64, voucherId int64) error {
	if _, err := s.VoucherUserRepository.CreateOne(
		s.VoucherRepository.GetDatabase().Db,
		table.VoucherUser.MutableColumns,
		entity.VoucherUser{
			FkVoucher: voucherId,
			FkUser:    userId,
		},
	); err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.NoDataFound {
				return err
			}
		}
	}
	return nil
}
