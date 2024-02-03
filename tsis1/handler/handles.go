package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-practice/tsis1/data"
	"net/http"
)

func GetClubs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.FootballClubs)
}

func GetClub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clubID, exists := params["club"]
	if !exists {
		http.Error(w, "Identification not specified ", http.StatusBadRequest)
		return
	}

	club, exists := data.FootballClubs[clubID]
	if !exists {
		http.Error(w, "Club not found! ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(club)

}
