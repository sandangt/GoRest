package middleware

import (
	"net/http"
	"strings"
	"GoRest/service"
	"GoRest/util"
	)

//var availablePublishers []interface{} = []interface{} { "facebook", "twitter", "snapchat", "pinterest", "linkedin" }

func CheckPublisher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		availablePublishers := service.ReadAllPublishers()
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
		archived := strings.ToLower(req.FormValue("archived"))
		if archived == "" || archived == "true" || archived == "false" {
			next.ServeHTTP(res, req)
		} else {
			data:= service.InvalidArchivedStatus()
			util.PackingSendingData(res, req, http.StatusBadRequest, &data)
		}	
	})
}

func CheckContinuousStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		isContinuous := strings.ToLower(req.FormValue("isContinuous"))
		if isContinuous == "" || isContinuous == "true" || isContinuous == "false" {
			next.ServeHTTP(res, req)
		} else {
			data:= service.InvalidContinuousStatus()
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
