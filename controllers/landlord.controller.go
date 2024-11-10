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
	var landlordReq requests.StoreLandlordRequest
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

	landlord := requests.ConvertLandlordStoreRequestToModel(landlordReq)
	/*
		THIS IS FOR PREVENTING EMAIL DUPLICATION

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
	*/
	newLandlord, errDb := repository.StoreLandlord(landlord)
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "CreateFailed",
			"data":    errDb.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    newLandlord,
	})
}

func GetLandlordByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	landlord, err := repository.FindLandlordByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "FindFailed",
			"data":    err.Error(),
		})
		return
	}

	if landlord.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "NoEntity",
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    landlord,
	})
}
