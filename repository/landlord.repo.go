package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreLandlord(landlord models.Landlord) (models.Landlord, error) {
	err := database.DB.Model(&models.Landlord{}).Create(&landlord).Error
	return landlord, err
}

func FindLandlordByEmail(Email string) (models.Landlord, error) {
	landlord := new(models.Landlord)
	err := database.DB.Model(&models.Landlord{}).Where("email=?", Email).Find(&landlord).Error

	return *landlord, err
}

func UpdateLandlord(ID string, landlord models.Landlord) (models.Landlord, error) {
	err := database.DB.Model(&models.Landlord{}).Where("id=?", ID).Updates(&landlord).Error
	return landlord, err
}
