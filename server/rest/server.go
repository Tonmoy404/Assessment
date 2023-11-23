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

	///////Brand Routes
	router.POST("/api/brand", server.createBrand)
	router.GET("/api/brand/:id", server.getBrand)
	router.GET("/api/brand/all", server.getBrands)
	router.DELETE("/api/brand/:id", server.deleteBrand)
	router.PATCH("/api/brand/:id", server.updateBrand)

	////Supplier
	router.POST("/api/supplier", server.createSupplier)
	router.GET("/api/supplier/:id", server.getSupplier)
	router.GET("/api/supplier/all", server.getSuppliers)
	router.DELETE("/api/supplier/:id", server.deleteSupplier)
	router.PATCH("/api/supplier/:id", server.updateSupplier)

	/////product routes
	router.POST("/api/product", server.createProduct)
	router.GET("/api/product/:id", server.getProduct)
	router.GET("/api/product/all", server.getProducts)
	router.DELETE("/api/product/:id", server.deleteProduct)
	router.PATCH("/api/product/:id", server.updateProduct)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}
