package routes

import (
	"noticepros/config/app_config"
	store_controllers "noticepros/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	landlordRoute := route.Group("landlord")

	landlordRoute.POST("/", store_controllers.StoreLandlord)
	// ROUTE STATIC
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
}
