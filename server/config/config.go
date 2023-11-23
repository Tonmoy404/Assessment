package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var appOnce = sync.Once{}
var dbOnce = sync.Once{}

type DB struct {
	Host   string `mapstructure:"DB_HOST"`
	Port   string `mapstructure:"DB_PORT"`
	User   string `mapstructure:"DB_USER"`
	Pass   string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}
type Application struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

var appConfig *Application
var dbConfig *DB

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

	dbConfig = &DB{
		Host:   viper.GetString("DB_HOST"),
		Port:   viper.GetString("DB_PORT"),
		DBName: viper.GetString("DB_NAME"),
		User:   viper.GetString("DB_USER"),
		Pass:   viper.GetString("DB_PASS"),
	}
}

func GetApp() *Application {
	appOnce.Do(func() {
		loadApp()
	})
	return appConfig
}

func GetDB() *DB {
	dbOnce.Do(func() {
		loadTable()
	})
	return dbConfig
}
