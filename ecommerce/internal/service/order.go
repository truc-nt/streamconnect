package service

import (
	"ecommerce/api/model"
	entity "ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"
	internalModel "ecommerce/internal/model"
	"ecommerce/internal/repository"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IOrderService interface {
	CreateOrderWithCartItems(createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest) error
}

type OrderService struct {
	OrderRepository                              repository.IOrderRepository
	OrderItemRepository                          repository.IOrderItemRepository
	OrderItemLivestreamExternalVariantRepository repository.IOrderItemLivestreamExternalVariantRepository
	CartItemRepository                           repository.ICartItemRepository
	ExternalOrderRepository                      repository.IExternalOrderRepository

	EcommerceService map[int16]IEcommerceService
}

func NewOrderService(
	orderRepository repository.IOrderRepository,
	orderItemRepository repository.IOrderItemRepository,
	orderItemLivestreamExternalVariantRepository repository.IOrderItemLivestreamExternalVariantRepository,
	cartItemRepository repository.ICartItemRepository,
	externalOrderRepository repository.IExternalOrderRepository,
	ecommerceService map[int16]IEcommerceService,
) IOrderService {
	return &OrderService{
		OrderRepository:     orderRepository,
		OrderItemRepository: orderItemRepository,
		OrderItemLivestreamExternalVariantRepository: orderItemLivestreamExternalVariantRepository,
		CartItemRepository:                           cartItemRepository,
		ExternalOrderRepository:                      externalOrderRepository,

		EcommerceService: ecommerceService,
	}
}

func (s *OrderService) CreateOrderWithCartItems(createOrderWithCartItemsRequest *model.CreateOrderWithCartItemsRequest) error {
	var execWinthinTransaction = func(db qrm.Queryable) (interface{}, error) {
		newOrderData := entity.Order{
			FkUser: createOrderWithCartItemsRequest.IDUser,
		}
		newOrder, err := s.OrderRepository.CreateOne(
			db,
			postgres.ColumnList{
				table.Order.FkUser,
			},
			newOrderData,
		)
		if err != nil {
			return nil, err
		}

		cartItemsByEcommerce, err := s.CartItemRepository.GetCartItemsByIds(db, createOrderWithCartItemsRequest.IDCartItems)
		if err != nil {
			return nil, err
		}

		externalOrderItemsByExternalShopIdAndEcommerce := make(map[int16]map[int64][]*internalModel.ExternalOrderItem)

		for _, cartItemByEcommerce := range *cartItemsByEcommerce {
			for _, cartItemsByShop := range cartItemByEcommerce.CartItemsGroupByShop {
				for _, cartItem := range cartItemsByShop.CartItems {
					newOrderItemData := entity.OrderItem{
						FkOrder:           newOrder.IDOrder,
						FkExternalVariant: cartItem.IDExternalVariant,
						Quantity:          cartItem.Quantity,
					}

					newOrderItem, err := s.OrderItemRepository.CreateOne(
						db,
						postgres.ColumnList{
							table.OrderItem.FkOrder,
							table.OrderItem.FkExternalVariant,
							table.OrderItem.Quantity,
						},
						newOrderItemData,
					)
					if err != nil {
						return nil, err
					}

					if _, err := s.OrderItemLivestreamExternalVariantRepository.CreateOne(
						db,
						postgres.ColumnList{
							table.OrderItemLivestreamExternalVariant.FkOrderItem,
							table.OrderItemLivestreamExternalVariant.Fk,
						},
						entity.OrderItemLivestreamExternalVariant{
							FkOrderItem: newOrderItem.IDOrderItem,
							Fk:          cartItem.IDLivestreamExternalVariant,
						},
					); err != nil {
						return nil, err
					}

					if _, ok := externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce]; !ok {
						externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce] = make(map[int64][]*internalModel.ExternalOrderItem)
					}

					if _, ok := externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce][cartItem.IDExternalShop]; !ok {
						externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce][cartItem.IDExternalShop] = []*internalModel.ExternalOrderItem{}
					}

					externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce][cartItem.IDExternalShop] = append(
						externalOrderItemsByExternalShopIdAndEcommerce[cartItemByEcommerce.IDEcommerce][cartItem.IDExternalShop],
						&internalModel.ExternalOrderItem{
							ExternalIdMapping: cartItem.ExternalIDMapping,
							Quantity:          int64(cartItem.Quantity),
						},
					)
				}
			}

		}

		if err := s.CartItemRepository.UpdateCartItemsToInactiveByIds(db, createOrderWithCartItemsRequest.IDCartItems); err != nil {
			return nil, err
		}

		for ecommerceId, externalOrderItemsGroupByExternalShopId := range externalOrderItemsByExternalShopIdAndEcommerce {
			for externalShopId, externalOrderItems := range externalOrderItemsGroupByExternalShopId {
				externalOrderIdMapping, err := s.EcommerceService[ecommerceId].CreateOrder(externalShopId, externalOrderItems)
				if err != nil {
					return nil, err
				}

				if _, err := s.ExternalOrderRepository.CreateOne(
					db,
					postgres.ColumnList{
						table.ExternalOrder.FkOrder,
						table.ExternalOrder.ExternalOrderIDMapping,
					},
					entity.ExternalOrder{
						FkOrder:                newOrder.IDOrder,
						ExternalOrderIDMapping: externalOrderIdMapping,
					},
				); err != nil {
					return nil, err
				}
			}
		}
		return nil, nil
	}

	if _, err := s.OrderRepository.ExecWithinTransaction(execWinthinTransaction); err != nil {
		return err
	}
	return nil
}
