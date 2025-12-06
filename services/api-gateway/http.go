package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/contracts"

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

	response := contracts.APIResponse{}

	if err := writeJSON(w, http.StatusOK, response); err != nil {
		log.Println("Failed to write JSON response:", err)
	}
}
