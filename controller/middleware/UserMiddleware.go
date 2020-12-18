package middleware

import (
	"net/http"
	
	"github.com/gorilla/mux"
	
	"GoRest/entity"
	"GoRest/service"
	"GoRest/util"
)

func CheckUserExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Check if user with given identityID exists
		// entity.User{} : empty struct
		if (service.ReadUserByIdentityID(mux.Vars(req)["identityID"]) == entity.User{}) {
			data:= service.UserNotFound()
			util.PackingSendingData(res, req, http.StatusNotFound, &data)
		} else {
			next.ServeHTTP(res, req)
		}	
	})
}	

