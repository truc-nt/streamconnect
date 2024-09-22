package cmd

import (
	"ecommerce/api/handler"
	"ecommerce/internal/adapter"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"fmt"

	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewShopHandler,
	handler.NewProductHandler,
	handler.NewVariantHandler,
	handler.NewShopifyHandler,
	handler.NewExternalShopHandler,
	handler.NewExternalVariantHandler,
	handler.NewLivestreamHandler,
	handler.NewLivestreamProductHandler,

	handler.NewCartHandler,
	handler.NewOrderHandler,
	handler.NewVoucherHandler,

	handler.ProvideHandlers,
)

var ServicesSet = wire.NewSet(
	service.NewUserService,
	service.NewShopService,

	service.NewShopifyService,
	service.ProvideEcommerceServices,

	service.NewProductService,
	service.NewVariantService,
	service.NewExternalShopService,
	service.NewExternalShopAuthService,
	service.NewExternalVariantService,
	service.NewLivestreamService,
	service.NewLivestreamProductService,
	service.NewCartService,
	service.NewUserAddressService,

	service.NewOrderService,
	service.NewVoucherService,
)

var RepositoriesSet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewUserAddressRepository,
	repository.NewShopRepository,

	repository.NewProductRepository,
	repository.NewVariantRepository,
	repository.NewImageVariantRepository,
	repository.NewExternalShopRepository,
	repository.NewExternalShopShopifyAuthRepository,
	repository.NewExternalVariantRepository,

	repository.NewLivestreamRepository,
	repository.NewLivestreamProductRepository,
	repository.NewLivestreamExternalVariantRepository,

	repository.NewCartItemRepository,
	repository.NewCartItemLivestreamExternalVariantRepository,

	repository.NewOrderRepository,
	repository.NewOrderItemRepository,
	repository.NewOrderItemLivestreamExternalVariantRepository,
	repository.NewExternalOrderRepository,
	repository.NewCartRepository,
	repository.NewLivestreamProductFollowerRepository,
	repository.NewVoucherRepository,
	repository.NewVoucherUserRepository,
)

var AdapterSet = wire.NewSet(
	adapter.NewShopifyAdapter,
	adapter.NewVideoSdkAdapter,
)

func runServer() {
	//var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	svr := initServer()
	svr.Start()

	defer func() {
		svr.Stop()
	}()
}
