package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/appointments/infra/gorm/entities"
	repositories2 "ismaeldf/golang-gobarber/modules/appointments/repositories"
	services2 "ismaeldf/golang-gobarber/modules/appointments/services"
	"ismaeldf/golang-gobarber/modules/users/infra/http/middlewares"
	"net/http"
)

var appointmentRepository = repositories2.AppointmentsRepository{}

func createAppointment(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	appointment := entities.Appointment{}
	_ = json.Unmarshal(b, &appointment)

	appointmentService := services2.NewCreateAppointmentService(&appointmentRepository)

	appointmentCreated, err := appointmentService.Execute(appointment)
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

	subRouter := mux.NewRouter().PathPrefix(path).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("", createAppointment).Methods("POST")
	subRouter.HandleFunc("", listAppointment).Methods("GET")

	router.PathPrefix(path).Handler(negroni.New(
		middlewares.EnsureAuthenticated(),
		negroni.Wrap(subRouter),
	))
}
