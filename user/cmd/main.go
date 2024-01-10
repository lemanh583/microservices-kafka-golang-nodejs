package main

import (
	"log"
	"user-services/config"
	"user-services/database"
	"user-services/di"
	"user-services/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading godotenv: ", err)
	}
	config.InitConfig()
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Error auto migrate: ", err)
	}
	r := gin.Default()
	d := di.Initialized(db)
	router.InitRouter(r, d)
	r.Run("127.0.0.1:" + config.Cfg.ServerPort)
}
