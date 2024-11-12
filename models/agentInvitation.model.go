package models

import (
	"noticepros/utils"

	"gorm.io/gorm"
)

type AgentInvitation struct {
	ID             string `gorm:"primaryKey"`
	OID            string `gorm:"not null"`
	Owner          User   `gorm:"foreignKey:OID"`
	AgentID        string `gorm:"not null"`
	Agent          User   `gorm:"foreignKey:AgentID"`
	Role           string // 'manager', 'attorney'
	InvitationDate string
	Status         string // 'Pending', 'Accepted', 'Rejected'
}

func (agentInvitation *AgentInvitation) BeforeCreate(tx *gorm.DB) (err error) {
	agentInvitation.ID = utils.RandomString(10)
	return
}
