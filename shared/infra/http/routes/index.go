package routes

import (
	"github.com/gorilla/mux"
	"ismaeldf/golang-gobarber/modules/appointments/infra/http/routes"
	routes2 "ismaeldf/golang-gobarber/modules/users/infra/http/routes"
	services2 "ismaeldf/golang-gobarber/modules/users/services"
	"net/http"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcome)

	router.Handle("/files/{rest}",
		http.StripPrefix("/files/", http.FileServer(http.Dir(services2.FileDirectory))))

	routes2.UsersRouter(router)
	routes2.SessionsRouter(router)
	routes.AppointmentsRouter(router)

	return router
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome API GO BARBER"))
}
