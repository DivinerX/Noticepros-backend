package models

import (
	"noticepros/utils"
	"time"

	"gorm.io/gorm"
)

type Attorney struct {
	ID            string `gorm:"type:string;primaryKey"`
	Password      string `gorm:"column:pwd"`
	FirstName     string `gorm:"column:fname"`
	LastName      string `gorm:"column:lname"`
	BusinessName  string `gorm:"column:bname"`
	Address       string `gorm:"column:address"`
	Unit          string `gorm:"column:unit"`
	City          string `gorm:"column:cty"`
	County        string `gorm:"column:cnty"`
	State         string `gorm:"column:st"`
	ZipCode       string `gorm:"column:zip"`
	TelePhone     string `gorm:"column:tel"`
	TelePhoneCell string `gorm:"column:telcel"`
	TelePhoneFax  string `gorm:"column:telfax"`
	Email1        string `gorm:"column:eml1"`
	Email2        string `gorm:"column:eml2"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (attorney *Attorney) BeforeCreate(tx *gorm.DB) (err error) {
	attorney.ID = utils.RandomString(10)
	attorney.Password = utils.RandomString(10)
	return
}
