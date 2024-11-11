package models

import (
	"time"

	"noticepros/utils"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:string;primaryKey"`
	Password      string `gorm:"column:pwd"`
	Role          string
	FirstName     string `gorm:"column:fname"`
	LastName      string `gorm:"column:lname"`
	BusinessName  string `gorm:"column:bname"`
	Address       string
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = utils.RandomString(10)
	user.Password = utils.RandomString(10)
	return
}
