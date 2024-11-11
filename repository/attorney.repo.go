package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func FindAttorneyByEmail(Email string) (models.Attorney, error) {
	attorney := new(models.Attorney)
	err := database.DB.Model(&models.Attorney{}).Where("email=?", Email).Find(&attorney).Error

	return *attorney, err
}
