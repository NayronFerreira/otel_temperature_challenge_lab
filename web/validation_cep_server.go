package web

import (
	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/NayronFerreira/otel_temperature_challenge_lab/service/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CEPValidator(config config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/", api.NewAPI(config).ValidatorCEP)
	return router
}
