package models

import (
	"math/rand"
	"time"

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

func randomString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (landlord *Landlord) BeforeCreate(tx *gorm.DB) (err error) {
	landlord.ID = randomString(10)
	landlord.Password = randomString(10)
	return
}
