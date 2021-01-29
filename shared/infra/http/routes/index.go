package routes

import (
	"github.com/gorilla/mux"
	appointmentsRoutes "ismaeldf/golang-gobarber/modules/appointments/infra/http/routes"
	userRoutes "ismaeldf/golang-gobarber/modules/users/infra/http/routes"
	services "ismaeldf/golang-gobarber/modules/users/services"
	"net/http"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcome)

	router.Handle("/files/{rest}",
		http.StripPrefix("/files/", http.FileServer(http.Dir(services.FileDirectory))))

	userRoutes.UsersRouter(router)
	userRoutes.SessionsRouter(router)
	appointmentsRoutes.AppointmentsRouter(router)

	return router
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome API GO BARBER"))
}
