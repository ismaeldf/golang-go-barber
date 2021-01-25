package repositories

import (
	"gorm.io/gorm"
	"ismaeldf.melo/golang/go-barber/models"
	"time"
)

type AppointmentRepositoryDTO struct {
	ProviderId string
	Date time.Time
}

type AppointmentsRepository struct {
	DB *gorm.DB
}

func (r *AppointmentsRepository) FindByDate(date time.Time) models.Appointment {
	var appointment models.Appointment
	r.DB.Where("date = ?", date).Find(&appointment)

	return appointment
}

func (r *AppointmentsRepository) Create(data AppointmentRepositoryDTO) (*models.Appointment, error) {
	appointment := models.NewAppointment(data.ProviderId, data.Date)

	err := r.DB.Create(appointment).Error
	if err != nil{
		return nil, err
	}

	return appointment, nil
}

func (r *AppointmentsRepository) All() []models.Appointment {
	var appointments[] models.Appointment

	r.DB.Find(&appointments)

	return appointments
}