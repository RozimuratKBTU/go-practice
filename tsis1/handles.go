package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(footballClubs)
}

func getClub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clubID, exists := params["club"]
	if !exists {
		http.Error(w, "Identification not specified ", http.StatusBadRequest)
		return
	}

	club, exists := footballClubs[clubID]
	if !exists {
		http.Error(w, "Club not found! ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(club)

}
