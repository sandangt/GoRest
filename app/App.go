package app

import (
	"github.com/gorilla/mux"
	"GoRest/controller"
	"GoRest/controller/middleware"
)


func Routing(router *mux.Router) {
	router.HandleFunc("/users/{identityID}/line_items", controller.ReadLineItemsByUserIdentityID).Methods("GET")
	router.Use(middleware.CheckUserExist)
	router.Use(middleware.CheckPublisher)
	router.Use(middleware.CheckArchived)
}
