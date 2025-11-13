package httpx

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON payload with the provided status code.
func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload == nil {
		return
	}
	_ = json.NewEncoder(w).Encode(payload)
}

// Error writes a standard error envelope.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]any{
		"error":   message,
		"status":  status,
		"success": false,
	})
}
