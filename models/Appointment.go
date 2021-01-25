package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	ID         string    `json:"id" gorm:"type:uuid;primary_key"`
	ProviderId string    `json:"provider_id"`
	Provider   User      `gorm:"foreignKey:ProviderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Date       time.Time `json:"date" gorm:"notnull"`
	CreatedAt  time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewAppointment(providerId string, date time.Time) *Appointment {
	return &Appointment{
		ID:         uuid.NewV4().String(),
		ProviderId: providerId,
		Date:       date,
	}
}
