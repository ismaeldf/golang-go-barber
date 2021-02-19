package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/repositories"
	"ismaeldf/golang-gobarber/modules/users/providers/HashProvider/implementations"
	implementations2 "ismaeldf/golang-gobarber/modules/users/providers/TokenProvider/implementations"
	"ismaeldf/golang-gobarber/modules/users/services"
	"net/http"
)

var sessionsControllerRepository = repositories.UsersRepository{}
var sessionsControllerHashProvider = implementations.BCryptHashProvider{}
var sessionsControllerTokenProvider = implementations2.JwtTokenProvider{}

type SessionsController struct{}

type requestDTO struct {
	Email string
	Password string
}

func (c SessionsController) SessionsCreate(w http.ResponseWriter, r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	body := requestDTO{}
	_ = json.Unmarshal(b, &body)

	authenticateUserService := services.NewAuthenticateUserService(
		&sessionsControllerRepository,
		&sessionsControllerHashProvider,
		&sessionsControllerTokenProvider,
	)

	userAuthenticated, err := authenticateUserService.Execute(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	js, _ := json.Marshal(userAuthenticated)

	return js
}
