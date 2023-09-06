package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucasRejanio/grade-avarage-api/models"
	"github.com/LucasRejanio/grade-avarage-api/utils"
)

func CalculateAverageHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.NewLogger()

	var request models.AverageRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	average := utils.CalculateAverage(request.Numbers)

	response := models.AverageResponse{
		Average: average,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		logger.Println("Server response status: 200")
	}
}
