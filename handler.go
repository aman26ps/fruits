package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fruits API is running"))
}

func getAllFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fruits)
}

func getFruitByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, fruit := range fruits {
		if fruit.ID == id {
			json.NewEncoder(w).Encode(fruit)
			return
		}
	}

	http.Error(w, "Fruit not found", http.StatusNotFound)
}

func addFruit(w http.ResponseWriter, r *http.Request) {
	var fruit Fruit
	json.NewDecoder(r.Body).Decode(&fruit)

	fruit.ID = getNextID()
	fruits = append(fruits, fruit)
	saveFruits()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fruit)
}
