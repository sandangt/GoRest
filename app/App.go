package app

import (
	"github.com/gorilla/mux"
	"GoRest/controller"
)


func Routing(router *mux.Router) {

	router.HandleFunc("/users/{identityID}/line_items", controller.ReadLineItemsByUserIdentityID).Methods("GET")
}
