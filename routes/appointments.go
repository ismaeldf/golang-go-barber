package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"ismaeldf.melo/golang/go-barber/services"
	"net/http"
)

var appointmentRepository = repositories.AppointmentsRepository{}

func createAppointment(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	appointment := models.Appointment{}
	err := json.Unmarshal(b, &appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	appointmentService := services.NewCreateAppointmentService(&appointmentRepository)

	appointmentCreated, err := appointmentService.Execute(services.Request{
		Provider: appointment.Provider,
		Date: appointment.Date,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(appointmentCreated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func listAppointment(w http.ResponseWriter, r *http.Request) {
	appointments := appointmentRepository.All()
	fmt.Print(appointments)
	js, err := json.Marshal(appointments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func AppointmentsRouter(router *mux.Router) {
	path := "/appointments"
	router.HandleFunc(path, createAppointment).Methods("POST")
	router.HandleFunc(path, listAppointment).Methods("GET")
}
