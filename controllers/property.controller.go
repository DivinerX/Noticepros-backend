package controllers

import (
	"fmt"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func StoreProperty(ctx *gin.Context) {
	var propertyReq requests.PropertyRequest
	if err := ctx.ShouldBindJSON(&propertyReq); err != nil {
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

	user, _ := ctx.Get("user")
	userData := user.(models.User)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    token,
	})
}
