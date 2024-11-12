package models

import (
	"time"

	"noticepros/utils"

	"gorm.io/gorm"
)

type Particular struct {
	ID          string `gorm:"primaryKey"`
	RentFrom    string
	RentThrough string
	Dollars     uint
	Cents       uint
	Written     string
	PayToFirst  string
	PayToLast   string
	Telephone   string
	Address     string
	City        string `gorm:"column:cty"`
	Unit        string `gorm:"column:unit"`
	State       string `gorm:"column:st"`
	ZipCode     string `gorm:"column:zip"`
	County      string `gorm:"column:cnty"`
	OpenHours   string
	OpenDays    string
	Property    Property `gorm:"foreignKey:PID"`
	PID         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (particular *Particular) BeforeCreate(tx *gorm.DB) (err error) {
	particular.ID = utils.RandomString(10)
	return
}
