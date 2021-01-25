package main

import (
	"ismaeldf.melo/golang/go-barber/database"
	"ismaeldf.melo/golang/go-barber/routes"
	"log"
	"net/http"
)

func main(){
	database.CreateConnectionDB()

	router := routes.LoadRoutes()

	err := http.ListenAndServe(":3333", router)
	if err != nil{
		log.Fatalf("Error starting server.")
		panic(err)
	}
}
