package helpers

import (
	"encoding/json"
	"net/http"
)

//StatusResponse is a model for JSON response to client
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// WriteJSON writes json value
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j, _ := json.Marshal(v)

	w.Write(j)
}

// WriteErrorJSON writes a json error
func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, StatusResponse{
		Status:  "error",
		Message: message,
	})
}
