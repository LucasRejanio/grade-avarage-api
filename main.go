package main

import (
	"log"
	"net/http"

	"github.com/LucasRejanio/grade-avarage-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calculate-average", handlers.CalculateAverageHandler).Methods("POST")

	port := "8000"
	log.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
