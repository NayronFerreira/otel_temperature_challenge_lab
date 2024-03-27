package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/NayronFerreira/otel_temperature_challenge_lab/service/cep_validator"
)

type Request struct {
	CEP string `json:"cep"`
}

func (a API) ValidatorCEP(w http.ResponseWriter, r *http.Request) {
	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if service.IsInvalidCEP(req.CEP) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprintf(w, "CEP %s is valid", req.CEP)
}
