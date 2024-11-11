package middleware

import (
	"fmt"
	"net/http"
	"noticepros/config/app_config"
	"noticepros/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "TokenInvalid",
			"data":    "TokenMissing",
		})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(app_config.SECRET_KEY), nil
	})

	if err != nil {
		fmt.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "TokenExpired",
				"data":    claims["exp"].(float64),
			})
			return
		}

		userID, ok := claims["sub"].(string)
		println(userID)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "TokenInvalid",
				"data":    "UserIDInvalid",
			})
			return
		}
		user, err := repository.FindUserByID(userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "TokenInvalid",
				"data":    err,
			})
			return
		}

		if user.ID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "TokenInvalid",
				"data":    "NoUser",
			})
			return
		}
		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "TokenInvalid",
			"data":    "TokenInvalid",
		})
		return
	}
	ctx.Next()
}
