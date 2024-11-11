package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func StoreUser(user models.User) (models.User, error) {
	err := database.DB.Model(&models.User{}).Create(&user).Error
	return user, err
}

func FindUserByEmail(Email string) (models.User, error) {
	user := new(models.User)
	err := database.DB.Model(&models.User{}).Where("eml1=?", Email).Find(&user).Error

	return *user, err
}

func FindUserByID(ID string) (models.User, error) {
	user := new(models.User)
	err := database.DB.Model(&models.User{}).Where("id=?", ID).Find(&user).Error

	return *user, err
}

func UpdateUser(ID string, user models.User) (models.User, error) {
	err := database.DB.Model(&models.User{}).Where("id=?", ID).Updates(&user).Error
	return user, err
}
