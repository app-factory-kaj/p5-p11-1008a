package handlers

import (
	"encoding/json"
	"net/http"
)

type greetingResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// Greeting handles GET /hello?name=X.
func Greeting(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "name query parameter is required"})
		return
	}
	writeJSON(w, http.StatusOK, greetingResponse{Message: "Hello, " + name + "!"})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
