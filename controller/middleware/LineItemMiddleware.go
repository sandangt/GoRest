package middleware

import (
	"net/http"
	"GoRest/service"
	"GoRest/util"
	)

var availablePublishers []interface{} = []interface{} { "facebook", "twitter", "snapchat", "pinterest", "linkedin" }

func CheckPublisher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		publisher := req.FormValue("publisher")
		if publisher == "" || findInSlice(availablePublishers, publisher) {
			next.ServeHTTP(res, req)
		} else {
			data:= service.InvalidPublisher()
			util.PackingSendingData(res, req, http.StatusBadRequest, &data)
		}	
	})
}

func CheckArchived(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		archived := req.FormValue("archived")
		if archived == "" || archived == "true" || archived == "false" {
			next.ServeHTTP(res, req)
		} else {
			data:= service.InvalidArchivedStatus()
			util.PackingSendingData(res, req, http.StatusBadRequest, &data)
		}	
	})
}
	
func findInSlice(slice []interface{}, target interface{}) bool {
	for _, v := range slice {
		if (target == v) {
			return true
		}
	}
	return false
}
