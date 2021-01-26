package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", welcome)

	router.Handle("/files/{rest}",
		http.StripPrefix("/files/", http.FileServer(http.Dir("tmp/"))))

	UsersRouter(router)
	SessionsRouter(router)
	AppointmentsRouter(router)

	return router
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome API GO BARBER"))
}
