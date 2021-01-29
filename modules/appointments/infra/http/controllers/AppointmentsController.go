package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/repositories"
	"ismaeldf/golang-gobarber/modules/appointments/services"
	"net/http"
)

var appointmentRepository = repositories.AppointmentsRepository{}

type AppointmentsController struct{}

func (c AppointmentsController) AppointmentCreate(w http.ResponseWriter, r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	appointment := entities.Appointment{}
	_ = json.Unmarshal(b, &appointment)

	appointmentService := services.NewCreateAppointmentService(&appointmentRepository)

	appointmentCreated, err := appointmentService.Execute(appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	js, _ := json.Marshal(appointmentCreated)

	return js
}

func (c AppointmentsController) AppointmentGet(w http.ResponseWriter, r *http.Request) []byte {
	appointments := appointmentRepository.All()

	js, _ := json.Marshal(appointments)

	return js
}