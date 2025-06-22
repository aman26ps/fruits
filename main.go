package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	loadFruits() // load fruits from file at startup

	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")
	router.HandleFunc("/fruits", getAllFruits).Methods("GET")
	router.HandleFunc("/fruits/{id}", getFruitByID).Methods("GET")
	router.HandleFunc("/fruits", addFruit).Methods("POST")

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
