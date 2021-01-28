package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	repositories2 "ismaeldf/golang-gobarber/modules/appointments/repositories"
	"time"
)

type createAppointmentService struct {
	appointmentRepository *repositories2.AppointmentsRepository
}

func NewCreateAppointmentService(repository *repositories2.AppointmentsRepository) *createAppointmentService {
	return &createAppointmentService{repository}
}

func (s *createAppointmentService) Execute(appointment entities.Appointment) (*entities.Appointment, error) {
	appointment.Date = startOfHour(appointment.Date)

	find := s.appointmentRepository.FindByDate(appointment.Date)
	if find.Id != "" {
		return nil, errors.New("This appointment is already booked")
	}

	appointmentCreated, err := s.appointmentRepository.Create(appointment)
	if err != nil {
		return nil, err
	}

	return appointmentCreated, nil
}

func startOfHour(date time.Time) time.Time {
	loc, _ := time.LoadLocation("UTC")

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		date.Hour(),
		0,
		0,
		0,
		loc,
	)
}
