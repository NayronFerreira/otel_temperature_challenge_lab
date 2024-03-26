package config

import (
	"os"
	"testing"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	os.Setenv("WEB_SERVER_PORT", ":8181")
	os.Setenv("VIACEP_HOST_API", "https://viacep.com.br/ws")
	os.Setenv("WEATHER_HOST_API", "http://api.weatherapi.com/v1")
	os.Setenv("WEATHER_API_KEY", "c5b200eed4fd417e9bc202030241003")
	os.Setenv("TIMEOUT_SECONDS", "5")

	config, err := config.LoadConfig(".")

	assert.NoError(t, err)
	assert.Equal(t, ":8181", config.WebValidationCEPServerPort)
	assert.Equal(t, "https://viacep.com.br/ws", config.ViaCepHost)
	assert.Equal(t, "http://api.weatherapi.com/v1", config.WeatherHost)
	assert.Equal(t, "c5b200eed4fd417e9bc202030241003", config.WeatherKey)
	assert.Equal(t, "5", config.TimeoutSeconds)
}
