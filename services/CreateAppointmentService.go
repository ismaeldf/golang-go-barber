package services

import (
	"errors"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"time"
)

type createAppointmentService struct {
	appointmentRepository *repositories.AppointmentsRepository
}

func NewCreateAppointmentService(repository *repositories.AppointmentsRepository) *createAppointmentService {
	return &createAppointmentService{repository}
}

func (s *createAppointmentService) Execute(appointment models.Appointment) (*models.Appointment, error) {
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
