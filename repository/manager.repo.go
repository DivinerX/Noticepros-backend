package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func FindManagerByEmail(Email string) (models.Manager, error) {
	manager := new(models.Manager)
	err := database.DB.Model(&models.Manager{}).Where("email=?", Email).Find(&manager).Error

	return *manager, err
}
