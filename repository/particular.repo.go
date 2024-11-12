package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreParticular(particular models.Particular) (models.Particular, error) {
	err := database.DB.Model(&models.Particular{}).Create(&particular).Error
	return particular, err
}

func GetAllParticular() ([]models.Particular, error) {
	var particular []models.Particular
	err := database.DB.Model(&models.Particular{}).Preload("Property").Find(&particular).Error
	return particular, err
}
