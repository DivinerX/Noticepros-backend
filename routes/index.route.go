package routes

import (
	"noticepros/config/app_config"
	"noticepros/controllers"
	"noticepros/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app.Group("api")

	route.POST("/login", controllers.Login)

	userRoute := route.Group("user")
	userRoute.POST("/", controllers.StoreUser)
	userRoute.GET("/me", middleware.RequireAuth, controllers.GetUserInfo)

	propertyRoute := route.Group("property")
	propertyRoute.POST("/", middleware.RequireAuth, controllers.StoreProperty)
	propertyRoute.GET("/", controllers.GetAllProjects)

	tenantRoute := route.Group("tenant")
	tenantRoute.POST("/", middleware.RequireAuth, controllers.StoreTenant)
	tenantRoute.GET("/", controllers.GetAllTenants)
	// ROUTE STATIC
	route.Static(app_config.STATIC_ROUTE, app_config.STATIC_DIR)
}
