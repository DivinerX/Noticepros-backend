package routes

import (
	"noticepros/config/app_config"
	"noticepros/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	route.POST("/login", controllers.Login)

	landlordRoute := route.Group("landlord")
	landlordRoute.POST("/", controllers.StoreLandlord)
	landlordRoute.GET("/:id", controllers.GetLandlordByID)

	// ROUTE STATIC
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
}
