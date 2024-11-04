package bootstrap

import (
	"log"
	configs "noticepros/config"
	"noticepros/config/app_config"
	"noticepros/config/cors_config"
	"noticepros/config/validate_config"
	"noticepros/database"
	"noticepros/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", validate_config.PhoneValidator)
	}
	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(app_config.APP_PORT)
}
