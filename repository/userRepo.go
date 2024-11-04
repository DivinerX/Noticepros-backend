package repository

import (
	"noticepros/database"
	"noticepros/models"
)

func GetUserByTID(id string) (models.User, error) {
	var user models.User
	errDB := database.DB.Table("users").Where("t_id=?", id).First(&user).Error

	return user, errDB
}
