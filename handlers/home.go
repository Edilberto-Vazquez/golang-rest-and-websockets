package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{Message: "Welcome to Platzi Go", Status: true})
	}
}
