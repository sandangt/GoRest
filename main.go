package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	
	"GoRest/app"
	"GoRest/repository"
	"GoRest/configuration"
)

func main() {
	// Init config
	configInfo := configuration.ParseConfiguration()
	repository.CreateDriver(configInfo)
	defer repository.CloseDriver()
	
	// Init router
	router := mux.NewRouter()

	// Route handles & endpoints
	app.Routing(router)

	// Start server
	log.Fatal(http.ListenAndServe(":9062", router))
}

