package repositories

import (
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	"time"
)

type IAppointmentsRepository interface {
	FindByDate(date time.Time) entities.Appointment
	Create(data entities.Appointment) (*entities.Appointment, error)
	All() []entities.Appointment
}