package server

import (
	"log"
	"net/http"

	"github.com/NayronFerreira/microservice-ratelimiter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	mux := http.NewServeMux()
	mux.Handle("/check", handler)
	mux.Handle("/metrics", promhttp.Handler()) // prometheus metrics

	return &Server{
		Server: http.Server{
			Addr:    ":" + cfg.WebPort,
			Handler: mux,
		},
	}
}

func (s *Server) Start() {
	log.Printf("Rate Limit Middleware running on port %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
