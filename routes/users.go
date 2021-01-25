package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
	"ismaeldf.melo/golang/go-barber/services"
	"net/http"
)

var usersRepository = repositories.UsersRepository{}

func createUser(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	user := models.User{}
	_ = json.Unmarshal(b, &user)

	createUserService := services.NewCreateUserService(&usersRepository)

	userCreated, err := createUserService.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, _ := json.Marshal(userCreated)

	w.Write(js)
}

func UsersRouter(router *mux.Router){
	path := "/users"
	router.HandleFunc(path, createUser).Methods("POST")
}
