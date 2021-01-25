package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"ismaeldf.melo/golang/go-barber/services"
	"net/http"
)

type requestDTO struct {
	Email string
	Password string
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	body := requestDTO{}
	_ = json.Unmarshal(b, &body)

	authenticateUserService := services.NewAuthenticateUserService(&usersRepository)

	userAuthenticated, err := authenticateUserService.Execute(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, _ := json.Marshal(userAuthenticated)

	w.Write(js)
}

func SessionsRouter(router *mux.Router){
	path := "/sessions"
	router.HandleFunc(path, authenticate).Methods("POST")
}
