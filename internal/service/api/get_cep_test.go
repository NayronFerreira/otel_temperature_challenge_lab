package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/stretchr/testify/assert"
)

func TestGetLocationByCEP(t *testing.T) {
	// Cria um servidor HTTP de teste
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"localidade": "São Paulo", "uf": "SP"}`))
	}))
	defer ts.Close()

	// Configura a API para usar o URL do servidor de teste como o host ViaCep
	api := API{
		Config: config.Config{
			ViaCepHost:     ts.URL,
			TimeoutSeconds: "30",
		},
	}

	// Chama a função GetLocationByCEP
	locality, UF, err := api.GetLocationByCEP("01001000")
	if err != nil {
		t.Fatal(err)
	}

	// Verifica se a função GetLocationByCEP retornou os valores corretos
	assert.Equal(t, "São Paulo", locality)
	assert.Equal(t, "SP", UF)
	assert.Nil(t, err)
}
