package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureTypesByCEP(t *testing.T) {
	// Cria um servidor HTTP de teste para o ViaCep
	viaCepServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"localidade": "São Paulo", "uf": "SP"}`))
	}))
	defer viaCepServer.Close()

	// Cria um servidor HTTP de teste para o WeatherHost
	weatherHostServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"main": {"temp": 293.15}}`)) // 293.15K é 20°C
	}))
	defer weatherHostServer.Close()

	// Configura a API para usar o URL dos servidores de teste como o host ViaCep e WeatherHost
	api := API{
		Config: config.Config{
			ViaCepHost:     viaCepServer.URL,
			WeatherHost:    weatherHostServer.URL,
			TimeoutSeconds: "30",
		},
	}

	// Cria um ResponseWriter de teste
	w := httptest.NewRecorder()

	// Cria uma Request de teste
	r := httptest.NewRequest(http.MethodGet, "/temperature/01001000", nil)

	// Chama a função GetTemperatureTypesByCEP
	api.GetTemperatureTypesByCEP(w, r)

	// Verifica se a função GetTemperatureTypesByCEP retornou a resposta correta
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"temp_C":0,"temp_F":32,"temp_K":273.15}`+"\n", w.Body.String())
}
