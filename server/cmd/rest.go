package cmd

import (
	"fmt"
	"log"

	"github.com/Tonmoy404/Assessment/cache"
	"github.com/Tonmoy404/Assessment/config"
	"github.com/Tonmoy404/Assessment/repo"
	"github.com/Tonmoy404/Assessment/rest"
	"github.com/Tonmoy404/Assessment/service"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

const ERROR_TABLE = "error_tables"

func serveRest() {
	appConfig := config.GetApp()
	dbConfig := config.GetDB()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s username=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	brandRepo := repo.NewBrandRepo(db)
	productRepo := repo.NewProductRepo(db)
	categoryRepo := repo.NewCategoryRepo(db)
	supplierRepo := repo.NewSupplierRepo(db)
	errRepo := repo.NewErrorRepo(db, ERROR_TABLE)

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
