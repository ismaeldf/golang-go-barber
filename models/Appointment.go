package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Appointment struct {
	ID       string    `json:"id"`
	Provider string       `json:"provider"`
	Date     time.Time `json:"date"`
}

func NewAppointment(provider string, date time.Time) *Appointment{
	return &Appointment{
		ID: uuid.NewV4().String(),
		Provider : provider,
		Date: date,
	}
}