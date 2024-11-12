package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreTenant(tenant models.Tenant) (models.Tenant, error) {
	err := database.DB.Model(&models.Tenant{}).Create(&tenant).Error
	return tenant, err
}

func GetAllTenants() ([]models.Tenant, error) {
	var tenant []models.Tenant
	err := database.DB.Model(&models.Tenant{}).Preload("Property").Find(&tenant).Error
	return tenant, err
}
