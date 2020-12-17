package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	
	"github.com/gorilla/mux"
	"GoRest/service"
)

func ReadLineItemsByUserIdentityID(w http.ResponseWriter, r *http.Request) {
	/**
	 * Request params: 
	 * - id
	 * - name
	 * - publisher
	 * - pacingType
	 * - Archived
	 * - isContinous
	 * - CreatorCompanyName
	 * - BrandCompanyName
	 * - BrandName
	 * - InitiativeName
	 */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if (r.URL.Query()["test"] != nil) {
		fmt.Println(r.URL.Query()["test"][0])
	} else {
		fmt.Println("No")
	}
	data:= service.ReadLineItemsByUserIdentityID(mux.Vars(r)["identityID"])
	json.NewEncoder(w).Encode(data)
}
