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

func StoreTenant(ctx *gin.Context) {
	var tenantReq requests.TenantRequest
	if err := ctx.ShouldBindJSON(&tenantReq); err != nil {
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

	var tenants []models.Tenant
	for i := 0; i < len(tenantReq); i++ {
		tenant := models.Tenant{
			FirstName:     tenantReq[i].FirstName,
			LastName:      tenantReq[i].LastName,
			TelePhone:     tenantReq[i].TelePhone,
			TelePhoneCell: tenantReq[i].TelePhoneCell,
			TelePhoneFax:  tenantReq[i].TelePhoneFax,
			Email1:        tenantReq[i].Email1,
			Email2:        tenantReq[i].Email2,
			PID:           tenantReq[i].PID,
			Index:         uint(i + 1),
		}

		var err error
		tenant, err = repository.StoreTenant(tenant)

		if err != nil {
			println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "DBError",
				"data":    err,
			})
			return // Optional: Exit if storing a tenant fails
		}

		tenants = append(tenants, tenant) // Append the tenant to the slice
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    tenants,
	})
}

func GetAllTenants(ctx *gin.Context) {
	tenants, err := repository.GetAllTenants()

	if err != nil {
		println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    tenants,
	})
}
