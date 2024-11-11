package controllers

import (
	"fmt"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Login(ctx *gin.Context) {
	var loginReq requests.LoginRuquest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
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

	switch loginReq.Type {
	case 0:
		landlord, err := repository.FindLandlordByEmail(loginReq.Email)
		if err != nil {
			println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalError",
				"data":    err,
			})
			return
		}
		if landlord.Password != loginReq.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "InvalidPassword",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"data":    landlord,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    loginReq,
	})
}
