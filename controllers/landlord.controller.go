package store_controllers

import (
	"fmt"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func StoreLandlord(ctx *gin.Context) {
	var landlordReq requests.LandlordRequest
	if err := ctx.ShouldBindJSON(&landlordReq); err != nil {
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

	landlord := requests.ConvertLandlordRequestToModel(landlordReq)

	existLandlord, err := repository.FindLandlordByEmail(landlordReq.Email)

	if err != nil {
		println(err)
	}

	if existLandlord.ID != "" {
		ctx.AbortWithStatusJSON(http.StatusInsufficientStorage, gin.H{
			"message": "EmailAlreadyUse",
		})
		return
	}

	newLandlord, errDb := repository.StoreLandlord(landlord)
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "CreateFailed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "CreateSuccess",
		"data":    newLandlord,
	})
}
