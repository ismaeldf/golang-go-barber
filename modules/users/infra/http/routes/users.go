package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"ismaeldf/golang-gobarber/modules/users/infra/http/controllers"
	"ismaeldf/golang-gobarber/modules/users/infra/http/middlewares"
	"net/http"
)

var usersController = controllers.UsersController{}

func createUser(w http.ResponseWriter, r *http.Request) {
	js := usersController.UsersCreate(w, r)
	w.Write(js)
}

func updateAvatar(w http.ResponseWriter, r *http.Request) {
	js := usersController.UsersUpdateAvatar(w, r)
	w.Write(js)
}

func UsersRouter(router *mux.Router){
	path := "/users"

	router.HandleFunc(path, createUser).Methods("POST")

	subRouter := mux.NewRouter().PathPrefix(path).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/avatar", updateAvatar).Methods("PATCH")

	router.PathPrefix(path).Handler(negroni.New(
		middlewares.EnsureAuthenticated(),
		negroni.Wrap(subRouter),
	))
}
