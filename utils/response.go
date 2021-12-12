package utils

import (
	"encoding/json"
	"net/http"
)

func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	SetResponseHeaders(w)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "message": err.Error()})
}

func SucessResponse(w http.ResponseWriter) {
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func SucessResponseWithData(w http.ResponseWriter, data interface{}) {
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
