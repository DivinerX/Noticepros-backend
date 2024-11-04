package bootstrap

import (
	"log"
	configs "noticepros/config"
	"noticepros/config/app_config"
	"noticepros/config/cors_config"
	"noticepros/database"
	"noticepros/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootStrapApp() {

	// LOAD .env FILE
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIT CONFIG
	configs.InitConfig()

	// DATABASE CONNECTION
	database.ConnectDatabase()

	// INIT GIN ENGINE
	app := gin.Default()

	// CORS
	app.Use(cors_config.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(app_config.APP_PORT)
}
