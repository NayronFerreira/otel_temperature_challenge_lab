package controller

import (
	"github.com/NayronFerreira/microservice-input/internal/infra/web/handlers"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	router  chi.Router
	Handler *handlers.Handler
}

func NewController(
	router chi.Router,
	Handler *handlers.Handler,
) *Controller {
	return &Controller{
		router:  router,
		Handler: Handler,
	}
}

func (wc *Controller) Route() {
	wc.router.Route("/", func(r chi.Router) {
		r.Post("/", wc.Handler.GetTemperatures)
	})
}
