package controllers

import (
	"encoding/json"
	"io/ioutil"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/repositories"
	"ismaeldf/golang-gobarber/modules/users/infra/http/middlewares"
	"ismaeldf/golang-gobarber/modules/users/services"
	"net/http"
)

var usersControllerRepository = repositories.UsersRepository{}

type UsersController struct{}

func (c UsersController) UsersCreate(w http.ResponseWriter, r *http.Request) []byte {
	b, _ := ioutil.ReadAll(r.Body)

	user := entities.UserUnhide{}
	_ = json.Unmarshal(b, &user)

	createUserService := services.NewCreateUserService(&usersControllerRepository)

	userCreated, err := createUserService.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	js, _ := json.Marshal(userCreated)

	return js
}

func (c UsersController) UsersUpdateAvatar(w http.ResponseWriter, r *http.Request) []byte {
	//parse input max size 1 MB to 2 MB
	r.ParseMultipartForm(1 << 2)

	//retrieve file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	//write temporary file
	tempFile, err := ioutil.TempFile(services.FileDirectory, "avatar-*.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	tempFile.Write(fileBytes)

	userId := middlewares.GetUserIdContext(r)
	fileName := tempFile.Name()

	updateUserAvatarService := services.NewUpdateUserAvatarService((&usersControllerRepository))

	userUpdated, err := updateUserAvatarService.Execute(userId, fileName)

	js, _ := json.Marshal(userUpdated)

	return js
}