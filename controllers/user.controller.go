package controllers

import (
	"fmt"
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

func StoreUser(ctx *gin.Context) {
	var landlordReq requests.StoreUserRequest
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

	landlord := requests.ConvertUserStoreRequestToModel(landlordReq)

	// THIS IS FOR PREVENTING EMAIL DUPLICATION

	existUser, err := repository.FindUserByEmail(landlordReq.Email1)

	if err != nil {
		println(err)
	}

	if existUser.ID != "" {
		ctx.AbortWithStatusJSON(http.StatusInsufficientStorage, gin.H{
			"message": "EmailAlreadyUse",
		})
		return
	}

	newUser, errDb := repository.StoreUser(landlord)

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "CreateFailed",
			"data":    errDb.Error(),
		})
		return
	}

	claims := jwt.MapClaims{
		"sub": newUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, errToken := utils.GenerateToken(&claims)
	result := responses.UserResponse{
		ID:       newUser.ID,
		Password: newUser.Password,
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

func GetUserByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	user, err := repository.FindUserByID(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "FindFailed",
			"data":    err.Error(),
		})
		return
	}

	if user.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "NoEntity",
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    user,
	})
}
