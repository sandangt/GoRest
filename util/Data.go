package util

import (
	"net/http"
	"encoding/json"
)

func PackingSendingData(res http.ResponseWriter, req *http.Request, statusCode int, data *interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(*data)
}
