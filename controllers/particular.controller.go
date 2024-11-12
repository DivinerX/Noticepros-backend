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

func StoreParticular(ctx *gin.Context) {
	var particularReq requests.ParticularRequest
	if err := ctx.ShouldBindJSON(&particularReq); err != nil {
		println(err.Error())
		// if err.Error() != "" {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{
		// 		"message": "ValidateFailed",
		// 		"data":    err.Error(),
		// 	})
		// 	return
		// }
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

	particular := models.Particular{
		RentFrom:    particularReq.RentFrom,
		RentThrough: particularReq.RentThrough,
		Dollars:     particularReq.Dollars,
		Cents:       particularReq.Cents,
		Written:     particularReq.Written,
		PayToFirst:  particularReq.PayToFirst,
		PayToLast:   particularReq.PayToLast,
		Telephone:   particularReq.Telephone,
		Address:     particularReq.Address,
		City:        particularReq.City,
		Unit:        particularReq.Unit,
		State:       particularReq.State,
		ZipCode:     particularReq.ZipCode,
		County:      particularReq.County,
		OpenHours:   particularReq.OpenHours,
		OpenDays:    particularReq.OpenDays,
		PID:         particularReq.PID,
	}

	newParticular, err := repository.StoreParticular(particular)
	if err != nil {
		println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "DBError",
			"data":    err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    newParticular,
	})
}

func GetAllParticulars(ctx *gin.Context) {
	properties, err := repository.GetAllParticular()

	if err != nil {
		println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    properties,
	})
}
