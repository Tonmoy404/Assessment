package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}
var tableOnce = sync.Once{}
var saltOnce = sync.Once{}

type Table struct {
	BrandTableName        string `mapstructure:"BRAND_TABLE_NAME"`
	CategoryTableName     string `mapstructure:"CATEGORY_TABLE_NAME"`
	SupplierTableName     string `mapstructure:"SUPPLIER_TABLE_NAME"`
	ProductTableName      string `mapstructure:"PRODUCT_TABLE_NAME"`
	ProductStockTableName string `mapstructure:"PRODUCT_STOCK_TABLE_NAME"`
}

type Application struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

var appConfig *Application
var tableConfig *Table

func loadApp() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	appConfig = &Application{
		Host: viper.GetString("HOST"),
		Port: viper.GetString("PORT"),
	}
}

func loadTable() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file was not found, that's okay")
	}

	viper.AutomaticEnv()

	tableConfig = &Table{
		BrandTableName:        viper.GetString("BRAND_TABLE_NAME"),
		CategoryTableName:     viper.GetString("CATEGORY_TABLE_NAME"),
		SupplierTableName:     viper.GetString("SUPPLIER_TABLE_NAME"),
		ProductTableName:      viper.GetString("PRODUCT_TABLE_NAME"),
		ProductStockTableName: viper.GetString("PRODUCT_STOCK_TABLE_NAME"),
	}
}

func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})
	return appConfig
}

func GetTable() *Table {
	tableOnce.Do(func() {
		loadTable()
	})
	return tableConfig
}
