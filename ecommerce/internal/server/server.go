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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, user_id")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func NewServer(cfg *configs.Config, h *handler.Handlers) IServer {
	e := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	//e.Use(cors.New(config))
	e.Use(CORSMiddleware())

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
