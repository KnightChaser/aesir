package webRequestHandler

import (
	"encoding/json"
	"net/http"
)

// APIErrorJSONHandler is a function that returns a JSON response with an error message
func apiErrorJSONHandler(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": message,
	})
	http.Error(w, message, statusCode)
}
