package controllers

import (
	"fmt"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/models"
	"noticepros/repository"

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

	println(userData.ID)

	property := models.Property{
		Name:         propertyReq.Name,
		Address:      propertyReq.Address,
		City:         propertyReq.City,
		Unit:         propertyReq.Unit,
		State:        propertyReq.State,
		ZipCode:      propertyReq.ZipCode,
		County:       propertyReq.County,
		NumUnitTotal: propertyReq.NumUnitTotal,
		OID:          userData.ID,
	}

	newProperty, err := repository.StoreProperty(property)
	if err != nil {
		println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "DBError",
			"data":    err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    newProperty,
	})
}

func GetAllProjects(ctx *gin.Context) {
	properties, err := repository.GetAllProperty()

	if err != nil {
		println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    properties,
	})
}
