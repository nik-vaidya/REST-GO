package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rest_golang/controller"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/add", controller.Post).Methods("POST")
	r.HandleFunc("/retrieve", controller.Get).Methods("GET")
	r.HandleFunc("/update/{id}", controller.Put).Methods("PUT")
	r.HandleFunc("/delete/{id}", controller.Delete).Methods("DELETE")

	http.ListenAndServe(":8040", r)

}
