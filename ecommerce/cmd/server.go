package cmd

import (
	"ecommerce/api/handlers"
	"ecommerce/internal/adapters"
	"ecommerce/internal/database"
	"ecommerce/internal/repositories"
	"ecommerce/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	handlers.NewShopifyHandler,
	handlers.NewExternalShopHandler,
)

var ServicesSet = wire.NewSet(
	services.NewShopifyService,
	services.NewExternalShopService,
	services.NewExternalShopAuthService,
	services.ProvideEcommerceServices,
)

var RepositoriesSet = wire.NewSet(
	repositories.NewExternalShopRepository,
	repositories.NewShopifyExternalShopAuthRepository,
)

var AdapterSet = wire.NewSet(
	adapters.NewShopifyAdapter,
)

type Server struct {
	HttpServer *http.Server
	Engine     *gin.Engine

	Handlers *handlers.Handlers
	Services *services.Services

	PostgresDatabase *database.PostgresqlDatabase
}

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
