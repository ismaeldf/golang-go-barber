package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/repositories"
	"ismaeldf/golang-gobarber/modules/users/infra/http/middlewares"
	"ismaeldf/golang-gobarber/modules/users/providers/HashProvider/implementations"
	"ismaeldf/golang-gobarber/modules/users/services"
	implementations2 "ismaeldf/golang-gobarber/shared/container/providers/StorageProvider/implementations"
	"net/http"
)

var usersControllerRepository = repositories.UsersRepository{}
var usersControllerHashProvider = implementations.BCryptHashProvider{}
var usersControllerIStorageProvider = implementations2.DiskStorageProvider{}

type UsersController struct{}

func (c UsersController) UsersCreate(w http.ResponseWriter, r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	user := entities.UserUnhide{}
	_ = json.Unmarshal(b, &user)

	createUserService := services.NewCreateUserService(&usersControllerRepository, &usersControllerHashProvider)

	userCreated, err := createUserService.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	js, _ := json.Marshal(userCreated)

	return js
}

func (c UsersController) UsersUpdateAvatar(w http.ResponseWriter, r *http.Request) []byte {
	err := r.ParseMultipartForm(1 << 2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	userId := middlewares.GetUserIdContext(r)

	updateUserAvatarService := services.NewUpdateUserAvatarService(
		&usersControllerRepository,
		&usersControllerIStorageProvider,
	)

	userUpdated, err := updateUserAvatarService.Execute(userId, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	js, _ := json.Marshal(userUpdated)

	return js
}
