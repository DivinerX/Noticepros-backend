package requests

import "noticepros/models"

type PropertyRequest struct {
	Name         string
	Address      string `binding:"required"`
	City         string `binding:"required"`
	Unit         string
	State        string `binding:"required"`
	ZipCode      string `binding:"required"`
	County       string `binding:"required"`
	NumUnitTotal uint8
}

func ConvertPropertyRequestToModel(req PropertyRequest) models.Property {
	return models.Property{
		Name:         req.Name,
		Address:      req.Address,
		City:         req.City,
		Unit:         req.Unit,
		State:        req.State,
		ZipCode:      req.ZipCode,
		County:       req.County,
		NumUnitTotal: req.NumUnitTotal,
	}
}
