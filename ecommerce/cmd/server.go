package cmd

import (
	"ecommerce/api/handlers"
	"ecommerce/internal/database"
	"ecommerce/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
