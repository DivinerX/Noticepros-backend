package cors_config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsConfig(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Controll-Allow-Origin", "*")

	ctx.Writer.Header().Set("Access-Controll-Allow-Credential", "true")

	ctx.Writer.Header().Set("Access-Controll-Allow-Headers", "Content-Type, Content-Length")

	ctx.Writer.Header().Set("Access-Controll-Allow-Methods", "POST, GET, PUT, PATCH, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)

		return
	}

	ctx.Next()
}
