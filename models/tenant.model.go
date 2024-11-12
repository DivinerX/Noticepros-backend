package models

import (
	"time"

	"noticepros/utils"

	"gorm.io/gorm"
)

type Tenant struct {
	ID            string `gorm:"type:string;primaryKey"`
	Index         uint
	FirstName     string `gorm:"column:fname"`
	LastName      string `gorm:"column:lname"`
	TelePhone     string `gorm:"column:tel"`
	TelePhoneCell string `gorm:"column:telcel"`
	TelePhoneFax  string `gorm:"column:telfax"`
	Email1        string `gorm:"column:eml1"`
	Email2        string `gorm:"column:eml2"`
	SS            string
	DOB           string
	DL            string
	DLST          string
	Property      Property `gorm:"foreignKey:PID"`
	PID           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (tenant *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	tenant.ID = utils.RandomString(10)
	return
}
