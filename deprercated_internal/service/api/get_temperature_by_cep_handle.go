package api

import (
	"encoding/json"
	"net/http"

	service "github.com/NayronFerreira/otel_temperature_challenge_lab/internal/service/temperature"
	"github.com/go-chi/chi"
)

func (a API) GetTemperatureTypesByCEP(w http.ResponseWriter, r *http.Request) {

	cep := chi.URLParam(r, "cep")

	//Chamada Servico A
	if err := a.isValidCEP(cep); err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	locality, UF, err := a.GetLocationByCEP(cep)
	if err != nil {

		if err.Error() == "can not find zipcode" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	celcius, err := a.GetCelciusByLocality(locality, UF)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempConverted := service.GenerateTemperatureTypesByCelcius(celcius)
	tempConverted.City = locality

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(tempConverted); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
