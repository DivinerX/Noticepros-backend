package models

import (
	"time"

	"noticepros/utils"

	"gorm.io/gorm"
)

type Property struct {
	ID           string `gorm:"type:string;primaryKey"`
	Name         string `gorm:"column:name"`
	Address      string
	City         string `gorm:"column:cty"`
	Unit         string `gorm:"column:unit"`
	State        string `gorm:"column:st"`
	ZipCode      string `gorm:"column:zip"`
	County       string `gorm:"column:cnty"`
	NumUnitTotal uint8
	Owner        User `gorm:"foreignKey:OID"`
	OID          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (property *Property) BeforeCreate(tx *gorm.DB) (err error) {
	property.ID = utils.RandomString(10)
	return
}
