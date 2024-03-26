package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/stretchr/testify/assert"
)

func TestGetCelciusByLocality(t *testing.T) {
	// Cria um servidor HTTP de teste para o WeatherHost
	weatherHostServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"current": {"temp_c": 20.0}}`)) // 20°C
	}))
	defer weatherHostServer.Close()

	// Configura a API para usar o URL do servidor de teste como o WeatherHost
	api := API{
		Config: config.Config{
			WeatherHost:    weatherHostServer.URL,
			TimeoutSeconds: "30",
		},
	}

	// Chama a função GetCelciusByLocality
	celcius, err := api.GetCelciusByLocality("São Paulo", "SP")
	if err != nil {
		t.Fatal(err)
	}

	// Verifica se a função GetCelciusByLocality retornou o valor correto
	assert.Equal(t, 20.0, celcius)
	assert.Nil(t, err)
}
