package routes

import (
	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	UsersRouter(router)
	SessionsRouter(router)
	AppointmentsRouter(router)

	return router
}