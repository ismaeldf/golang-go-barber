package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Appointment struct {
	ID       string    `json:"id" gorm:"type:uuid;primary_key"`
    Provider string    `json:"provider" gorm:"type:uuid;notnull"`
	Date     time.Time `json:"date"`
}

func NewAppointment(provider string, date time.Time) *Appointment{
	return &Appointment{
		ID: uuid.NewV4().String(),
		Provider : provider,
		Date: date,
	}
}