package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"noticepros/dtos/requests"
	"noticepros/dtos/responses"
	"noticepros/repository"
	"noticepros/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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

	// THIS IS FOR PREVENTING EMAIL DUPLICATION

	existLandlord, err := repository.FindLandlordByEmail(landlordReq.Email1)

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
			"data":    errDb.Error(),
		})
		return
	}

	sub := responses.Sub{
		ID:   newLandlord.ID,
		Type: 1,
	}
	subData, err := json.Marshal(sub)
	if err != nil {
		log.Fatal(err)
	}
	claims := jwt.MapClaims{
		"sub": subData,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, errToken := utils.GenerateToken(&claims)
	result := responses.LandlordResponse{
		ID:       newLandlord.ID,
		Password: newLandlord.Password,
		Token:    token,
	}
	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "FailedGenerateToken",
			"data":    errToken,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    result,
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
