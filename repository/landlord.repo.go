package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreLandlord(landlord models.Landlord) (models.Landlord, error) {
	err := database.DB.Table("landlords").Create(&landlord).Error
	return landlord, err
}

func FindLandlordByEmail(Email string) (models.Landlord, error) {
	landlord := new(models.Landlord)
	err := database.DB.Table("landlords").Where("email=?", Email).Find(&landlord).Error

	return *landlord, err
}
