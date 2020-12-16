package controller

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
	"GoRest/service"
)

func ReadLineItemsByUserIdentityID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, code := service.ReadLineItemsByUserIdentityID(mux.Vars(r)["identityID"])
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
