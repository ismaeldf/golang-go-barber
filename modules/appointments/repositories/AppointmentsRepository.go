package repositories

import (
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	"ismaeldf/golang-gobarber/shared/infra/gorm"
	"time"
)

type AppointmentsRepository struct {}

func (r *AppointmentsRepository) FindByDate(date time.Time) entities.Appointment {
	var appointment entities.Appointment

	gorm.DB.Where("date = ?", date).Find(&appointment)

	return appointment
}

func (r *AppointmentsRepository) Create(data entities.Appointment) (*entities.Appointment, error) {
	appointment := entities.NewAppointment(data.ProviderId, data.Date)

	err := gorm.DB.Create(appointment).Error
	if err != nil{
		return nil, err
	}

	return appointment, nil
}

func (r *AppointmentsRepository) All() []entities.Appointment {
	var appointments[]entities.Appointment

	gorm.DB.Find(&appointments)

	return appointments
}