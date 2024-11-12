package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreProperty(property models.Property) (models.Property, error) {
	err := database.DB.Model(&models.Property{}).Create(&property).Error
	return property, err
}

func GetAllProperty() ([]models.Property, error) {
	var property []models.Property
	err := database.DB.Model(&models.Property{}).Preload("Owner").Find(&property).Error
	return property, err
}
