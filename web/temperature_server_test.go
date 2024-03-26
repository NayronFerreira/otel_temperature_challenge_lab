package web

import (
	"net/http"
	"testing"
	"time"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/stretchr/testify/assert"
)

func TestInitializeRoutes(t *testing.T) {
	// Configuração para o teste
	conf := config.Config{
		WebValidationCEPServerPort: ":8181",
		ViaCepHost:                 "https://viacep.com.br/ws",
		WeatherHost:                "http://api.weatherapi.com/v1",
		WeatherKey:                 "c5b200eed4fd417e9bc202030241003",
		TimeoutSeconds:             "30",
	}

	// Inicializa as rotas em uma nova goroutine
	go InitializeRoutes(conf)

	// Espera o servidor estar pronto
	time.Sleep(2 * time.Second)

	// Cria uma Request de teste
	req, err := http.NewRequest("GET", "http://localhost:8181/01001000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Faz a Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta tem o status code correto
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
