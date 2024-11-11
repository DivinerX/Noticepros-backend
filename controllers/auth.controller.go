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

	switch loginReq.Type {
	case 1:
		landlord, err := repository.FindLandlordByEmail(loginReq.Email)
		if err != nil {
			println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalError",
				"data":    err,
			})
			return
		}
		if landlord.ID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "IncorrectUser",
			})
			return
		}
		if landlord.Password != loginReq.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "IncorrectPassword",
			})
			return
		}
		sub := responses.Sub{
			ID:   landlord.ID,
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
		return
	case 2:
		manager, err := repository.FindManagerByEmail(loginReq.Email)
		if err != nil {
			println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalError",
				"data":    err,
			})
			return
		}
		if manager.ID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "IncorrectUser",
			})
			return
		}
		if manager.Password != loginReq.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "InvalidPassword",
			})
			return
		}
		sub := responses.Sub{
			ID:   manager.ID,
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
		return
	case 3:
		attorney, err := repository.FindAttorneyByEmail(loginReq.Email)
		if err != nil {
			println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "InternalError",
				"data":    err,
			})
			return
		}
		if attorney.ID == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "IncorrectUser",
			})
			return
		}
		if attorney.Password != loginReq.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "CredintialError",
				"data":    "InvalidPassword",
			})
			return
		}
		sub := responses.Sub{
			ID:   attorney.ID,
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
		return
	}
}
