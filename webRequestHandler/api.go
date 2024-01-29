package webRequestHandler

import (
	"encoding/json"
	"net/http"
)

func API(w http.ResponseWriter, r *http.Request) {
	// Create a map to represent the JSON data
	responseData := map[string]string{"data": "testing successful"}

	// Convert the map to JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
