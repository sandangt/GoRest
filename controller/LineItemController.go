package controller

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"GoRest/service"
)

func ReadLineItemsByUserIdentityID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data:= service.ReadLineItemsByUserIdentityID(mux.Vars(r)["identityID"])
	json.NewEncoder(w).Encode(data)
}
