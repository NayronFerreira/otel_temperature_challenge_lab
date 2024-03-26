package web

import (
	"fmt"
	"net/http"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
)

func InitializeRoutes(config config.Config) {
	// http.ListenAndServe(config.WebValidationCEPServerPort, ValidationCEPRoute(config))
	// fmt.Println("Server running on port", config.WebValidationCEPServerPort)
	http.ListenAndServe(config.WebTemperatureServerPort, TemperatureRoute(config))
	fmt.Println("Server running on port", config.WebTemperatureServerPort)
}
