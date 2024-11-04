package requests

import (
	"noticepros/models"
)

type LandlordRequest struct {
	FirstName     string `binding:"required"`
	LastName      string `binding:"required"`
	BusinessName  string
	StreetAddress string `binding:"required"`
	City          string `binding:"required"`
	Unit          string
	State         string `binding:"required"`
	Code          string `binding:"required"`
	County        string `binding:"required"`
	Phone         string `binding:"required,phone"`
	Email         string `binding:"required,email"`
}

func ConvertLandlordRequestToModel(req LandlordRequest) models.Landlord {
	return models.Landlord{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		BusinessName:  req.BusinessName,
		StreetAddress: req.StreetAddress,
		City:          req.City,
		Unit:          req.Unit,
		State:         req.State,
		Code:          req.Code,
		County:        req.County,
		Phone:         req.Phone,
		Email:         req.Email,
	}
}
