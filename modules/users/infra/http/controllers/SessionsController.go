package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/repositories"
	"ismaeldf/golang-gobarber/modules/users/services"
	"net/http"
)

var sessionsControllerRepository = repositories.UsersRepository{}

type SessionsController struct{}

type requestDTO struct {
	Email string
	Password string
}

func (c SessionsController) SessionsCreate(w http.ResponseWriter, r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	body := requestDTO{}
	_ = json.Unmarshal(b, &body)

	authenticateUserService := services.NewAuthenticateUserService(&sessionsControllerRepository)

	userAuthenticated, err := authenticateUserService.Execute(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	js, _ := json.Marshal(userAuthenticated)

	return js
}
