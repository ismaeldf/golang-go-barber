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
	appointment, err := entities.NewAppointment(data.ProviderId, data.Date)
	if err != nil{
		return nil, err
	}

	_ = gorm.DB.Create(appointment).Error

	return appointment, nil
}

func (r *AppointmentsRepository) All() []entities.Appointment {
	var appointments[]entities.Appointment

	gorm.DB.Find(&appointments)

	return appointments
}
