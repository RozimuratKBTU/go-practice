package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RouterMux() {
	router := mux.NewRouter()

	router.HandleFunc("/clubs", getClubs).Methods("GET")

	router.HandleFunc("/clubs/{club}", getClub).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
