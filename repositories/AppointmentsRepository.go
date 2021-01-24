package repositories

import (
	"ismaeldf.melo/golang/go-barber/models"
	"time"
)

type AppointmentRepositoryDTO struct {
	Provider string
	Date time.Time
}

type AppointmentsRepository struct {
	appointments []models.Appointment
}

func (r *AppointmentsRepository) FindByDate(date time.Time) models.Appointment {
	for _, appointment := range r.appointments {
		if appointment.Date == date {
			return appointment
		}
	}

	return models.Appointment{}
}

func (r *AppointmentsRepository) Create(data AppointmentRepositoryDTO) models.Appointment {
	appointment := models.NewAppointment(data.Provider, data.Date)

	r.appointments = append(r.appointments, *appointment)

	return *appointment
}

func (r *AppointmentsRepository) All() []models.Appointment {
	return r.appointments
}