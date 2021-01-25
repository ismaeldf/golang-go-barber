package routes

import "github.com/gorilla/mux"

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()

	AppointmentsRouter(router)
	UsersRouter(router)

	return router
}