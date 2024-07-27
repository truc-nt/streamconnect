//go:build wireinject
// +build wireinject

package cmd

import (
	"ecommerce/api/handlers"
	"ecommerce/internal/configs"
	"ecommerce/internal/database"
	"ecommerce/internal/server"
	"ecommerce/internal/services"

	"github.com/google/wire"
)

//go:generate wire -package=cmd
func initServer() server.IServer {
	wire.Build(
		configs.NewConfig,

		RepositoriesSet,

		AdapterSet,
		//adapters.NewEcommerceAdapter,

		//ExternalShopService,

		ServicesSet,
		services.NewServices,

		HandlersSet,
		handlers.NewHandlers,

		database.NewPostgresDatabase,

		server.NewServer,
	)
	return &server.Server{}
}
