package cmd

import (
	"database/sql"

	"github.com/Tonmoy404/Assessment/cache"
	"github.com/Tonmoy404/Assessment/config"
	"github.com/Tonmoy404/Assessment/repo"
	"github.com/Tonmoy404/Assessment/rest"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/go-redis/redis"
)

func servreRest() {
	appConfig := config.GetApp()
	tableConfig := config.GetTable()

	db, err := sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	brandRepo := repo.NewBrandRepo(db, tableConfig.BrandTableName)
	productRepo := repo.NewProductRepo(db, tableConfig.ProductTableName)
	categoryRepo := repo.NewCategoryRepo(db, tableConfig.CategoryTableName)
	supplierRepo := repo.NewSupplierRepo(db, tableConfig.SupplierTableName)
	errRepo := repo.NewErrorRepo(db, tableConfig.ErrorTableName)

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cache := cache.NewCache(redisClient)

	svc := service.NewService(brandRepo, categoryRepo, productRepo, supplierRepo, cache, errRepo)

	server, err := rest.NewServer(appConfig, svc)
	if err != nil {
		panic("Server can not start")
	}

	server.Start()

}
