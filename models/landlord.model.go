package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Landlord struct {
	ID            string `gorm:"type:uuid;primaryKey"`
	FirstName     string
	LastName      string
	BusinessName  string
	StreetAddress string
	City          string
	Unit          string
	State         string
	Code          string
	County        string
	Phone         string
	Email         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (landlord *Landlord) BeforeCreate(tx *gorm.DB) (err error) {
	landlord.ID = uuid.New().String()
	return
}
