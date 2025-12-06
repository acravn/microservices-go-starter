package main

import (
	"log"
	"net/http"

	"encoding/json"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}
	// TODO: Call trip service

	response, err := http.NewRequest("POST", "http://trip-service:8083/preview", nil)
	if err != nil {
		http.Error(w, "Failed to create request to trip service", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err)
		http.Error(w, "Failed to get preview from trip service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := writeJSON(w, http.StatusOK, resp.Body); err != nil {
		log.Println("Failed to write JSON response:", err)
	}
}
