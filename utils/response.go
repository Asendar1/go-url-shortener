package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message string `json:"message,omitempty"`
	Data	any		`json:"data,omitempty"`
	Error	string	`json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func JSONSuccess(w http.ResponseWriter, status int, data any) {
	WriteJSON(w, status, JSONResponse{Message: "success", Data: data})
}

func JSONError(w http.ResponseWriter, status int, err string) {
	WriteJSON(w, status, JSONResponse{Error: err})
}
