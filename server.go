package main

import (
	"ismaeldf.melo/golang/go-barber/routes"
	"net/http"
)

func main(){
	router := routes.LoadRoutes()
	http.ListenAndServe(":3333", router);
}
