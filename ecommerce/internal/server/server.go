package server

import (
	"fmt"
	"net/http"

	"ecommerce/api/handlers"
	"ecommerce/api/routers"
	"ecommerce/internal/configs"
	"ecommerce/internal/database"
	"ecommerce/internal/services"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start()
	Stop()
}
type Server struct {
	HttpServer *http.Server
	Engine     *gin.Engine

	Handlers *handlers.Handlers
	Services *services.Services

	PostgresDatabase *database.PostgresqlDatabase
}

func NewServer(cfg *configs.Config, h *handlers.Handlers, s *services.Services, pDb *database.PostgresqlDatabase) IServer {
	e := gin.Default()
	return &Server{
		HttpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: e,
		},
		Engine:   e,
		Handlers: h,
		Services: s,

		PostgresDatabase: pDb,
	}
}

func (s *Server) Start() {
	s.LoadRouter()
	if err := s.HttpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	s.PostgresDatabase.Db.Close()
	if err := s.HttpServer.Close(); err != nil {
		panic(err)
	}
}

func (s *Server) LoadRouter() {
	routers.LoadApiRouter(s.Engine, s.Handlers)
}
