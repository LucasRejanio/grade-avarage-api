package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/LucasRejanio/grade-avarage-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calculate-average", handlers.CalculateAverageHandler).Methods("POST")

	port := ":8000"
	log.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}