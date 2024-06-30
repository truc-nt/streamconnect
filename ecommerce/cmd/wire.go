//go:build wireinject
// +build wireinject

package cmd

import (
	"ecommerce/api/handlers"
	"ecommerce/internal/configs"
	"ecommerce/internal/database"
	"ecommerce/internal/repositories"
	"ecommerce/internal/server"
	"ecommerce/internal/services"

	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(
	handlers.NewShopifyHandler,
)

var ServicesSet = wire.NewSet(
	services.NewShopifyService,
)

var RepositoriesSet = wire.NewSet(
	repositories.NewShopifyRepository,
)

//go:generate wire -package=cmd
func initServer() server.IServer {
	wire.Build(
		configs.NewConfig,

		RepositoriesSet,

		ServicesSet,
		services.NewServices,

		HandlersSet,
		handlers.NewHandlers,

		database.NewPostgresDatabase,

		server.NewServer,
	)
	return &server.Server{}
}
