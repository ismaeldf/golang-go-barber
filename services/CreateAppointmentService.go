package services

import (
	"errors"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"time"
)

type Request struct {
	ProviderId string
	Date time.Time
}

type createAppointmentService struct {
	appointmentRepository *repositories.AppointmentsRepository
}

func NewCreateAppointmentService(repository *repositories.AppointmentsRepository) *createAppointmentService {
	return &createAppointmentService{repository}
}

func (s *createAppointmentService) Execute(data Request) (models.Appointment, error) {
	appointmentDate := startOfHour(data.Date)

	find := s.appointmentRepository.FindByDate(appointmentDate)
	if find.ID != "" {
		return models.Appointment{}, errors.New("This appointment is already booked")
	}

	appointment, err := s.appointmentRepository.Create(repositories.AppointmentRepositoryDTO{
		ProviderId: data.ProviderId,
		Date:     appointmentDate,
	})
	if err != nil {
		return models.Appointment{}, err
	}

	return *appointment, nil
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
