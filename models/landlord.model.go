package models

import (
	"time"

	"noticepros/utils"

	"gorm.io/gorm"
)

type Landlord struct {
	ID            string `gorm:"type:string;primaryKey"`
	Password      string `gorm:"column:pwd"`
	FirstName     string `gorm:"column:fname"`
	LastName      string `gorm:"column:lname"`
	BusinessName  string `gorm:"column:bname"`
	Address       string `gorm:"column:address"`
	City          string `gorm:"column:cty"`
	Unit          string `gorm:"column:unit"`
	State         string `gorm:"column:st"`
	ZipCode       string `gorm:"column:zip"`
	County        string `gorm:"column:cnty"`
	TelePhone     string `gorm:"column:tel"`
	TelePhoneCell string `gorm:"column:telcel"`
	TelePhoneFax  string `gorm:"column:telfax"`
	Email1        string `gorm:"column:eml1"`
	Email2        string `gorm:"column:eml2"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (landlord *Landlord) BeforeCreate(tx *gorm.DB) (err error) {
	landlord.ID = utils.RandomString(10)
	landlord.Password = utils.RandomString(10)
	return
}
