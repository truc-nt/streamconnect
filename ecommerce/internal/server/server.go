package server

import (
	"fmt"
	"net/http"

	"ecommerce/api/handler"
	"ecommerce/api/router"
	"ecommerce/internal/configs"
	"ecommerce/internal/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start()
	Stop()
}
type Server struct {
	HttpServer *http.Server
	Engine     *gin.Engine

	Handlers *handler.Handlers

	PostgresDatabase *database.PostgresqlDatabase
}

func NewServer(cfg *configs.Config, h *handler.Handlers) IServer {
	e := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	e.Use(cors.New(config))

	return &Server{
		HttpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: e,
		},
		Engine:   e,
		Handlers: h,
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
	router.LoadApiRouter(s.Engine, s.Handlers)
}
