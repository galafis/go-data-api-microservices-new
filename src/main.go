package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-data-api-microservices/src/models"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go Microservices API!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := []models.Data{
		{ID: "1", Name: "Item 1"},
		{ID: "2", Name: "Item 2"},
	}
	json.NewEncoder(w).Encode(data)
}
