package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"ismaeldf/golang-gobarber/modules/appointments/infra/http/controllers"
	"ismaeldf/golang-gobarber/modules/users/infra/http/middlewares"
	"ismaeldf/golang-gobarber/modules/users/providers/TokenProvider/implementations"
	"net/http"
)

var appointmentController = controllers.AppointmentsController{}
var appointmentControllerTokenProvider = implementations.JwtTokenProvider{}


func createAppointment(w http.ResponseWriter, r *http.Request) {
	js := appointmentController.AppointmentCreate(w, r)
	w.Write(js)
}

func listAppointment(w http.ResponseWriter, r *http.Request) {
	js := appointmentController.AppointmentGet(w, r)
	w.Write(js)
}

func AppointmentsRouter(router *mux.Router) {
	path := "/appointments"

	subRouter := mux.NewRouter().PathPrefix(path).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("", createAppointment).Methods("POST")
	subRouter.HandleFunc("", listAppointment).Methods("GET")

	router.PathPrefix(path).Handler(negroni.New(
		middlewares.EnsureAuthenticated(&appointmentControllerTokenProvider),
		negroni.Wrap(subRouter),
	))
}
