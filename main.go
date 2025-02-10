package main

import (
	"dataset/internal/adapter/http"
	"dataset/internal/infrastructure/config"
	"dataset/internal/infrastructure/persistence/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	// Load app config by env
	loadConfig()
	// init db
	db.InitializeDB()
}

func main() {
	// Init gin
	engine := gin.Default()
	http.RegisterRoutes(engine)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}

func loadConfig() {
	env := os.Getenv("APP_ENV")
	env = strings.TrimSpace(env)
	// Set up viper to read the yaml file
	viper.SetConfigName(fmt.Sprintf("application_%s.yml", env))
	viper.AddConfigPath("./internal/infrastructure/config")
	viper.SetConfigType("yaml")
	// Read the yaml file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config.AppConfig)
	if err != nil {
		panic(err)
	}
}
