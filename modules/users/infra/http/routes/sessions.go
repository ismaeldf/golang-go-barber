package routes

import (
	"github.com/gorilla/mux"
	"ismaeldf/golang-gobarber/modules/users/infra/http/controllers"
	"net/http"
)

var sessionsController = controllers.SessionsController{}

func authenticate(w http.ResponseWriter, r *http.Request) {
	js := sessionsController.SessionsCreate(w, r)
	w.Write(js)
}

func SessionsRouter(router *mux.Router){
	path := "/sessions"
	router.HandleFunc(path, authenticate).Methods("POST")
}
