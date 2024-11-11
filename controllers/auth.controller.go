package controllers

import (
	"fmt"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/models"
	"noticepros/repository"
	"noticepros/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	var loginReq requests.LoginRuquest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		println(err)
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("Field '%s' failed validation: %s", err.Field(), err.Tag()))
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ValidateFailed",
			"data":    validationErrors,
		})
		return
	}

	user, err := repository.FindUserByEmail(loginReq.Email)
	if err != nil {
		println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "InternalError",
			"data":    err,
		})
		return
	}
	if user.ID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "CredintialError",
			"data":    "IncorrectUser",
		})
		return
	}
	if user.Password != loginReq.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "CredintialError",
			"data":    "IncorrectPassword",
		})
		return
	}
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "FailedGenerateToken",
			"data":    errToken,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    token,
	})
}

func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userData := user.(models.User)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    userData,
	})
}
