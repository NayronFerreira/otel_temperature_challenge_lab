package service

import "github.com/NayronFerreira/otel_temperature_challenge_lab/internal/model/entity"

func GenerateTemperatureTypesByCelcius(celsiusTemperature float64) entity.TemperatureTypes {
	return entity.TemperatureTypes{
		Celsius:    celsiusTemperature,
		Fahrenheit: convertCelsiusToFahrenheit(celsiusTemperature),
		Kelvin:     convertCelsiusToKelvin(celsiusTemperature),
	}
}

func convertCelsiusToFahrenheit(celsiusTemperature float64) float64 {
	return celsiusTemperature*1.8 + 32
}

func convertCelsiusToKelvin(celsiusTemperature float64) float64 {
	return celsiusTemperature + 273.15
}
