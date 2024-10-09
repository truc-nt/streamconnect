package service

import (
	"ecommerce/api/model"
	"ecommerce/internal/constants"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	internalModel "ecommerce/internal/model"
	"ecommerce/internal/repository"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderService interface {
	CreateOrderWithCartItems(userId int64, createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest) error
	GetBuyOrders(userId int64) (interface{}, error)
	GetOrder(orderId int64) (interface{}, error)
	GetOrdersByShopId(shopId int64) (interface{}, error)
	CreateOrderWithLivestreamExtVariantId(userId int64, livestreamExtVariantId int64) error
}

type OrderService struct {
	OrderRepository                              repository.IOrderRepository
	VoucherRepository                            repository.IVoucherRepository
	OrderItemRepository                          repository.IOrderItemRepository
	OrderItemLivestreamExternalVariantRepository repository.IOrderItemLivestreamExternalVariantRepository
	CartItemRepository                           repository.ICartItemRepository
	ExternalOrderRepository                      repository.IExternalOrderRepository
	UserRepository                               repository.IUserRepository
	UserAddressRepository                        repository.IUserAddressRepository
	LivestreamExternalVariantRepository          repository.ILivestreamExternalVariantRepository

	EcommerceService map[int16]IEcommerceService
}

func NewOrderService(
	orderRepository repository.IOrderRepository,
	voucherRepository repository.IVoucherRepository,
	orderItemRepository repository.IOrderItemRepository,
	orderItemLivestreamExternalVariantRepository repository.IOrderItemLivestreamExternalVariantRepository,
	cartItemRepository repository.ICartItemRepository,
	externalOrderRepository repository.IExternalOrderRepository,
	userRepository repository.IUserRepository,
	userAddressRepository repository.IUserAddressRepository,
	livestreamExternalVariantRepository repository.ILivestreamExternalVariantRepository,
	ecommerceService map[int16]IEcommerceService,
) IOrderService {
	return &OrderService{
		OrderRepository:     orderRepository,
		VoucherRepository:   voucherRepository,
		OrderItemRepository: orderItemRepository,
		OrderItemLivestreamExternalVariantRepository: orderItemLivestreamExternalVariantRepository,
		CartItemRepository:                           cartItemRepository,
		ExternalOrderRepository:                      externalOrderRepository,
		UserRepository:                               userRepository,
		UserAddressRepository:                        userAddressRepository,
		LivestreamExternalVariantRepository:          livestreamExternalVariantRepository,

		EcommerceService: ecommerceService,
	}
}

func (s *OrderService) CreateOrderWithCartItems(userId int64, createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest) error {
	var execWinthinTransaction = func(db qrm.Queryable) (interface{}, error) {
		newOrderData := entity.Order{
			FkUser:           userId,
			FkShop:           createOrderWithCartItemsRequest.IDShop,
			FkPaymentMethod:  1,
			FkShippingMethod: 1,
		}
		newOrder, err := s.OrderRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Order.FkUser,
				table.Order.FkShop,
				table.Order.FkPaymentMethod,
				table.Order.FkShippingMethod,
			},
			newOrderData,
		)
		if err != nil {
			return nil, err
		}

		cartItemIds := []int64{}
		subTotal := 0.0
		externalOrders := make([]*internalModel.ExternalOrder, 0)
		for _, externalOrder := range createOrderWithCartItemsRequest.ExternalOrders {

			_externalOrder := &internalModel.ExternalOrder{
				ShippingFee:         externalOrder.ShippingFee,
				ShippingFeeDiscount: externalOrder.ShippingFeeDiscount,
				ExternalDiscount:    externalOrder.ExternalDiscount,
				ExternalOrderItems:  make([]*internalModel.ExternalOrderItem, 0),
			}

			cartItemIds = append(cartItemIds, externalOrder.CartItemIds...)

			for _, cartItemId := range externalOrder.CartItemIds {
				cartItemData, err := s.CartItemRepository.GetCartItemInfoById(db, cartItemId)
				if err != nil {
					return nil, err
				}

				newOrderItemData := entity.OrderItem{
					FkOrder:      newOrder.IDOrder,
					FkExtVariant: cartItemData.FkExtVariant,
					Quantity:     cartItemData.Quantity,
					UnitPrice:    cartItemData.Price,
					PaidPrice:    cartItemData.Price,
				}
				subTotal += cartItemData.Price * float64(cartItemData.Quantity)

				newOrderItem, err := s.OrderItemRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.OrderItem.FkOrder,
						table.OrderItem.FkExtVariant,
						table.OrderItem.Quantity,
						table.OrderItem.UnitPrice,
						table.OrderItem.PaidPrice,
					},
					newOrderItemData,
				)
				if err != nil {
					return nil, err
				}

				if _, err := s.OrderItemLivestreamExternalVariantRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.OrderItemLivestreamExtVariant.FkOrderItem,
						table.OrderItemLivestreamExtVariant.FkLivestreamExtVariant,
					},
					entity.OrderItemLivestreamExtVariant{
						FkOrderItem:            newOrderItem.IDOrderItem,
						FkLivestreamExtVariant: cartItemData.IDLivestreamExternalVariant,
					},
				); err != nil {
					return nil, err
				}

				_externalOrder.IDExternalShop = cartItemData.IDExternalShop
				_externalOrder.IDEcommerce = cartItemData.IDEcommerce
				_externalOrder.ExternalOrderItems = append(_externalOrder.ExternalOrderItems, &internalModel.ExternalOrderItem{
					ExternalIdMapping: cartItemData.ExternalIDMapping,
					Quantity:          int64(cartItemData.Quantity),
				})
			}

			vouchers, err := s.VoucherRepository.GetByIds(db, externalOrder.VoucherIds)
			if err != nil {
				return nil, err
			}

			internalDiscount := 0.0
			for _, voucher := range vouchers {
				if voucher.Type == constants.VoucherTypePercentage {
					internalDiscount += min(voucher.Discount*subTotal, *voucher.MaxDiscount)
				} else {
					internalDiscount += min(voucher.Discount, subTotal)
				}
			}

			_externalOrder.InternalDiscount = internalDiscount
			externalOrders = append(externalOrders, _externalOrder)
		}

		if err := s.CartItemRepository.UpdateCartItemsToInactiveByIds(db, cartItemIds); err != nil {
			return nil, err
		}

		user, err := s.UserRepository.GetById(db, userId)
		if err != nil {
			return nil, err
		}

		address, err := s.UserAddressRepository.GetById(db, createOrderWithCartItemsRequest.IDUserAddress)
		if err != nil {
			return nil, err
		}

		for _, externalOrder := range externalOrders {
			externalOrderIdMapping, err := s.EcommerceService[externalOrder.IDEcommerce].CreateOrder(user, address, externalOrder.IDExternalShop, externalOrder.ExternalOrderItems, externalOrder.InternalDiscount)
			if err != nil {
				return nil, err
			}
			if _, err := s.ExternalOrderRepository.CreateOne(
				db,
				postgres.ColumnList{
					table.ExtOrder.FkOrder,
					table.ExtOrder.FkExtShop,
					table.ExtOrder.ExtOrderIDMapping,
					table.ExtOrder.ShippingFee,
					table.ExtOrder.ShippingFeeDiscount,
					table.ExtOrder.InternalDiscount,
					table.ExtOrder.ExternalDiscount,
				},
				entity.ExtOrder{
					FkOrder:           newOrder.IDOrder,
					FkExtShop:         externalOrder.IDExternalShop,
					ExtOrderIDMapping: externalOrderIdMapping,
					InternalDiscount:  externalOrder.InternalDiscount,
				},
			); err != nil {
				return nil, err
			}
		}

		return nil, nil
	}

	if _, err := s.OrderRepository.ExecWithinTransaction(execWinthinTransaction); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) GetBuyOrders(userId int64) (interface{}, error) {
	return s.OrderItemRepository.GetByUserId(s.OrderRepository.GetDatabase().Db, userId)
}

func (s *OrderService) GetOrder(orderId int64) (interface{}, error) {
	return s.ExternalOrderRepository.GetExternalOrdersByOrderId(s.OrderRepository.GetDatabase().Db, orderId)
}

func (s *OrderService) GetOrdersByShopId(shopId int64) (interface{}, error) {
	//return s.OrderRepository.GetByShopId(s.OrderRepository.GetDatabase().Db, shopId)
	return nil, nil
}

func (s *OrderService) CreateOrderWithLivestreamExtVariantId(userId int64, livestreamExtVariantId int64) error {
	var execWinthinTransaction = func(db qrm.Queryable) (interface{}, error) {

		variant, err := s.LivestreamExternalVariantRepository.GetVariantById(s.LivestreamExternalVariantRepository.GetDatabase().Db, livestreamExtVariantId)
		if err != nil {
			return nil, err
		}

		newOrderData := entity.Order{
			FkUser:           userId,
			FkShop:           variant.IDShop,
			FkPaymentMethod:  1,
			FkShippingMethod: 1,
		}
		newOrder, err := s.OrderRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Order.FkUser,
				table.Order.FkShop,
				table.Order.FkPaymentMethod,
				table.Order.FkShippingMethod,
			},
			newOrderData,
		)
		if err != nil {
			return nil, err
		}

		newOrderItemData := entity.OrderItem{
			FkOrder:      newOrder.IDOrder,
			FkExtVariant: variant.IDExtVariant,
			Quantity:     1,
			UnitPrice:    variant.Price,
			PaidPrice:    variant.Price,
		}

		newOrderItem, err := s.OrderItemRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.OrderItem.FkOrder,
				table.OrderItem.FkExtVariant,
				table.OrderItem.Quantity,
				table.OrderItem.UnitPrice,
				table.OrderItem.PaidPrice,
			},
			newOrderItemData,
		)
		if err != nil {
			return nil, err
		}

		if _, err := s.OrderItemLivestreamExternalVariantRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.OrderItemLivestreamExtVariant.FkOrderItem,
				table.OrderItemLivestreamExtVariant.FkLivestreamExtVariant,
			},
			entity.OrderItemLivestreamExtVariant{
				FkOrderItem:            newOrderItem.IDOrderItem,
				FkLivestreamExtVariant: livestreamExtVariantId,
			},
		); err != nil {
			return nil, err
		}

		user, err := s.UserRepository.GetById(db, userId)
		if err != nil {
			return nil, err
		}

		address, err := s.UserAddressRepository.GetDefaultAddressByUserId(s.UserAddressRepository.GetDatabase().Db, userId)
		if err != nil {
			return nil, err
		}

		externalOrderIdMapping, err := s.EcommerceService[variant.FkEcommerce].CreateOrder(user, address, variant.IDExtShop, []*internalModel.ExternalOrderItem{
			{
				ExternalIdMapping: variant.ExtIDMapping,
				Quantity:          1,
			},
		}, 0)
		if err != nil {
			return nil, err
		}
		if _, err := s.ExternalOrderRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.ExtOrder.FkOrder,
				table.ExtOrder.FkExtShop,
				table.ExtOrder.ExtOrderIDMapping,
				table.ExtOrder.ShippingFee,
				table.ExtOrder.ShippingFeeDiscount,
				table.ExtOrder.InternalDiscount,
				table.ExtOrder.ExternalDiscount,
			},
			entity.ExtOrder{
				FkOrder:           newOrder.IDOrder,
				FkExtShop:         variant.IDExtShop,
				ExtOrderIDMapping: externalOrderIdMapping,
				InternalDiscount:  0,
			},
		); err != nil {
			return nil, err
		}

		return nil, nil
	}

	if _, err := s.OrderRepository.ExecWithinTransaction(execWinthinTransaction); err != nil {
		return err
	}
	return nil

}
