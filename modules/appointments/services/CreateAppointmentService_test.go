package services_test

import (
	"github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	_ "github.com/stretchr/testify/require"
	fakesApppointmentsRepository "ismaeldf/golang-gobarber/modules/appointments/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/appointments/services"
	"testing"
	"time"
)

func TestCreateAppointmentService_Execute(t *testing.T) {
	t.Run("should be able to create a new appointment", func(t *testing.T) {
		apppointmentsRepository := fakesApppointmentsRepository.FakeAppointmentsRepository{}

		appointmentService := services.NewCreateAppointmentService(&apppointmentsRepository)

		appointment, _ := entities.NewAppointment("111", time.Now())

		appointmentCreated, _ := appointmentService.Execute(*appointment)

		require.Equal(t, appointmentCreated.ProviderId, appointment.ProviderId)
	})

	t.Run("should not be able to create two appointments in the same time", func(t *testing.T) {
		apppointmentsRepository := fakesApppointmentsRepository.FakeAppointmentsRepository{}

		appointmentService := services.NewCreateAppointmentService(&apppointmentsRepository)

		appointment, _ := entities.NewAppointment("111", time.Now())

		_, _ = appointmentService.Execute(*appointment)
		_, err := appointmentService.Execute(*appointment)

		require.NotNil(t, err)
	})

}
