package models

import (
	"noticepros/utils"

	"gorm.io/gorm"
)

type PropertyAgent struct {
	ID         string `gorm:"primaryKey"`
	PropertyID string `gorm:"not null"`
	Property   Property
	AgentID    string `gorm:"not null"`
	Agent      User   `gorm:"foreignKey:AgentID"`
	StartDate  string
	EndDate    string
	Role       string // 'manager', 'attorney'
}

func (propertyAgent *PropertyAgent) BeforeCreate(tx *gorm.DB) (err error) {
	propertyAgent.ID = utils.RandomString(10)
	return
}
