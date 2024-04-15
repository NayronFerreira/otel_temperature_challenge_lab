package utils

import (
	"net/http"

	"github.com/NayronFerreira/microservice-ratelimiter/model/dto"
	"github.com/goccy/go-json"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func JsonResponse(w http.ResponseWriter, response dto.ResponseDTO) {
	res := Response{
		Message: response.Message,
		Success: response.Success,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(jsonResponse)
}
