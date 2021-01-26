package main

import (
	"ismaeldf/golang-gobarber/database"
	"ismaeldf/golang-gobarber/routes"
	"log"
	"net/http"
)

func main(){
	database.CreateConnectionDB()

	router := routes.LoadRoutes()


	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("tmp"))))

	err := http.ListenAndServe(":3333", router)
	if err != nil{
		log.Fatalf("Error starting server.")
		panic(err)
	}
}

