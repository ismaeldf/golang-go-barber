package main

import (
	"ismaeldf/golang-gobarber/shared/infra/gorm"
	"ismaeldf/golang-gobarber/shared/infra/http/routes"
	"log"
	"net/http"
)

func main(){
	gorm.CreateConnectionDB()

	router := routes.LoadRoutes()

	err := http.ListenAndServe(":3333", router)
	if err != nil{
		log.Fatalf("Error starting server.")
		panic(err)
	}
}

