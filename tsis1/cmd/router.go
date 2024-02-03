package main

import (
	"github.com/gorilla/mux"

	"go-practice/tsis1/handler"
	"net/http"
)

func RouterMux() {
	router := mux.NewRouter()

	router.HandleFunc("/clubs", handler.GetClubs).Methods("GET")

	router.HandleFunc("/clubs/{club}", handler.GetClub).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
