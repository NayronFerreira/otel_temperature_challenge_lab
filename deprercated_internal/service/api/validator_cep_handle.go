package api

import (
	"encoding/json"
	"net/http"

	model "github.com/NayronFerreira/otel_temperature_challenge_lab/internal/model/request"
	service "github.com/NayronFerreira/otel_temperature_challenge_lab/internal/service/cep_validator"
)

func (a API) ValidatorCEP(w http.ResponseWriter, r *http.Request) {
	var req model.ValidatorCEPReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if service.IsInvalidCEP(req.CEP) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)
}
