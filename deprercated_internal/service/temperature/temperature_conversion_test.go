package service

import (
	"testing"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/internal/model/entity"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTemperatureTypesByCelcius(t *testing.T) {
	// Chama a função GenerateTemperatureTypesByCelcius
	temperatureTypes := GenerateTemperatureTypesByCelcius(20.0)

	// Verifica se a função GenerateTemperatureTypesByCelcius retornou os valores corretos
	expected := entity.TemperatureTypes{
		Celsius:    20.0,
		Fahrenheit: 68.0,
		Kelvin:     293.15,
	}
	assert.Equal(t, expected, temperatureTypes)
}

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	// Chama a função convertCelsiusToFahrenheit
	fahrenheit := convertCelsiusToFahrenheit(20.0)

	// Verifica se a função convertCelsiusToFahrenheit retornou o valor correto
	assert.Equal(t, 68.0, fahrenheit)
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	// Chama a função convertCelsiusToKelvin
	kelvin := convertCelsiusToKelvin(20.0)

	// Verifica se a função convertCelsiusToKelvin retornou o valor correto
	assert.Equal(t, 293.15, kelvin)
}
