package repositories

import (
	"ismaeldf.melo/golang/go-barber/database"
	"ismaeldf.melo/golang/go-barber/models"
	"time"
)

type AppointmentsRepository struct {}

func (r *AppointmentsRepository) FindByDate(date time.Time) models.Appointment {
	var appointment models.Appointment

	database.DB.Where("date = ?", date).Find(&appointment)

	return appointment
}

func (r *AppointmentsRepository) Create(data models.Appointment) (*models.Appointment, error) {
	appointment := models.NewAppointment(data.ProviderId, data.Date)

	err := database.DB.Create(appointment).Error
	if err != nil{
		return nil, err
	}

	return appointment, nil
}

func (r *AppointmentsRepository) All() []models.Appointment {
	var appointments[] models.Appointment

	database.DB.Find(&appointments)

	return appointments
}