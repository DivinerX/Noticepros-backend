package requests

import (
	"noticepros/models"
)

type StoreUserRequest struct {
	FirstName     string `binding:"required"`
	LastName      string `binding:"required"`
	Role          string `binding:"required"`
	BusinessName  string
	Address       string `binding:"required"`
	City          string `binding:"required"`
	Unit          string
	State         string `binding:"required"`
	ZipCode       string `binding:"required"`
	County        string `binding:"required"`
	TelePhone     string `binding:"required,phone"`
	TelePhoneCell string `binding:"required,phone"`
	TelePhoneFax  string `binding:"required,phone"`
	Email1        string `binding:"required,email"`
	Email2        string `binding:"omitempty,email"`
}

func ConvertUserStoreRequestToModel(req StoreUserRequest) models.User {
	return models.User{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Role:          req.Role,
		BusinessName:  req.BusinessName,
		Address:       req.Address,
		City:          req.City,
		Unit:          req.Unit,
		State:         req.State,
		ZipCode:       req.ZipCode,
		County:        req.County,
		TelePhone:     req.TelePhone,
		TelePhoneCell: req.TelePhoneCell,
		TelePhoneFax:  req.TelePhoneFax,
		Email1:        req.Email1,
		Email2:        req.Email2,
	}
}
