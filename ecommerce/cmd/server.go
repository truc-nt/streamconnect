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
	handler.NewProductHandler,
	handler.NewVariantHandler,
	handler.NewShopifyHandler,
	handler.NewExternalShopHandler,
	handler.NewExternalProductHandler,
	handler.NewLivestreamHandler,

	handler.ProvideHandlers,
)

var ServicesSet = wire.NewSet(
	service.NewShopifyService,
	service.ProvideEcommerceServices,

	service.NewProductService,
	service.NewVariantService,
	service.NewExternalShopService,
	service.NewExternalShopAuthService,
	service.NewExternalProductService,
	service.NewLivestreamService,
)

var RepositoriesSet = wire.NewSet(
	repository.NewProductRepository,
	repository.NewVariantRepository,
	repository.NewExternalShopRepository,
	repository.NewExternalShopShopifyAuthRepository,
	repository.NewExternalVariantRepository,

	repository.NewLivestreamRepository,
	repository.NewLivestreamProductRepository,
	repository.NewLivestreamExternalVariantRepository,
)

var AdapterSet = wire.NewSet(
	adapter.NewShopifyAdapter,
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
