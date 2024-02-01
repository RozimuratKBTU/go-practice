package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// Добавьте другие поля по необходимости
}

var friends = map[string]Friend{
	"john":   {"John Doe", 25},
	"alice":  {"Alice Smith", 28},
	"bob":    {"Bob Johnson", 22},
	"emily":  {"Emily Davis", 30},
	"charly": {"Charly Brown", 26},
	// Добавьте других друзей по необходимости
}

func main() {
	router := mux.NewRouter()

	// Маршрут для получения списка всех друзей
	router.HandleFunc("/friends", getFriends).Methods("GET")

	// Маршрут для получения информации о конкретном друге
	router.HandleFunc("/friends/{name}", getFriend).Methods("GET")

	// Запуск сервера на порту 8080
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func getFriends(w http.ResponseWriter, r *http.Request) {
	friendList := make([]Friend, 0, len(friends))
	for _, friend := range friends {
		friendList = append(friendList, friend)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friendList)
}

func getFriend(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, exists := params["name"]
	if !exists {
		http.Error(w, "Имя друга не указано", http.StatusBadRequest)
		return
	}

	friend, exists := friends[name]
	if !exists {
		http.Error(w, "Друг не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friend)
}
