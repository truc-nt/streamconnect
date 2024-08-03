//go:build wireinject
// +build wireinject

package cmd

import (
	"ecommerce/internal/configs"
	"ecommerce/internal/database"
	"ecommerce/internal/server"

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

		HandlersSet,

		database.NewPostgresDatabase,

		server.NewServer,
	)
	return &server.Server{}
}
