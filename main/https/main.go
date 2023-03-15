package main

import (
	"github.com/jcrob2/calc-apps/handlers"
	"net/http"
)

func main() {

	println("creating router...")
	router := handlers.SetupRouter()

	println("starting server...")
	http.ListenAndServe("localhost:8080", router)

}
