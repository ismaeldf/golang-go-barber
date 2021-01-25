package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"ismaeldf.melo/golang/go-barber/database"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"ismaeldf.melo/golang/go-barber/services"
	"net/http"
)

var appointmentRepository = repositories.AppointmentsRepository{
	DB: database.CreateConnectionDB(),
}

func createAppointment(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	appointment := models.Appointment{}
	_ = json.Unmarshal(b, &appointment)

	appointmentService := services.NewCreateAppointmentService(&appointmentRepository)

	appointmentCreated, err := appointmentService.Execute(services.RequestCreateAppointment{
		ProviderId: appointment.ProviderId,
		Date: appointment.Date,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, _ := json.Marshal(appointmentCreated)

	w.Write(js)
}

func listAppointment(w http.ResponseWriter, r *http.Request) {
	appointments := appointmentRepository.All()

	js, _ := json.Marshal(appointments)

	w.Write(js)
}

func AppointmentsRouter(router *mux.Router) {
	path := "/appointments"
	router.HandleFunc(path, createAppointment).Methods("POST")
	router.HandleFunc(path, listAppointment).Methods("GET")
}
