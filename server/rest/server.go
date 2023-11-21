package rest

import (
	"fmt"

	"github.com/Tonmoy404/Assessment/config"
	"github.com/Tonmoy404/Assessment/logger"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	appConfig *config.Application
	svc       service.Service
}

func NewServer(appConfig *config.Application, svc service.Service) (*Server, error) {
	server := &Server{
		appConfig: appConfig,
		svc:       svc,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(corsMiddleware)
	router.Use(logger.ModifyContext)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}
