package web

import (
	"fmt"
	"net/http"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
)

func InitializeRoutes(config config.Config) {

	go http.ListenAndServe(config.WebValidationCEPServerPort, CEPValidator(config))
	fmt.Println("Server running on port", config.WebValidationCEPServerPort)

	http.ListenAndServe(config.WebTemperatureServerPort, TemperatureRoute(config))
	fmt.Println("Server running on port", config.WebTemperatureServerPort)
}
