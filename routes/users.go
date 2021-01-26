package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io/ioutil"
	"ismaeldf/golang-gobarber/middlewares"
	"ismaeldf/golang-gobarber/models"
	"ismaeldf/golang-gobarber/repositories"
	"ismaeldf/golang-gobarber/services"
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

func updateAvatar(w http.ResponseWriter, r *http.Request) {
	//parse input max size 1 MB to 2 MB
	r.ParseMultipartForm(1 << 2)

	//retrieve file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//write temporary file
	tempFile, err := ioutil.TempFile(services.FileDirectory, "avatar-*.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tempFile.Write(fileBytes)

	userId := middlewares.GetUserIdContext(r)
	fileName := tempFile.Name()

	updateUserAvatarService := services.NewUpdateUserAvatarService(&usersRepository)

	userUpdated, err := updateUserAvatarService.Execute(userId, fileName)

	js, _ := json.Marshal(userUpdated)

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
