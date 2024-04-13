package main

import (
	"context"
	"log"

	"github.com/NayronFerreira/microservice-ratelimiter/config"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/database"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/handler"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/middleware"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/server"
	"github.com/NayronFerreira/microservice-ratelimiter/ratelimiter"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	rateLimitedHandler := middleware.RateLimitMiddleware(handler.DummyHandler(), SetupRateLimiter(cfg))

	server.NewServer(cfg, rateLimitedHandler).Start()

}

func SetupRateLimiter(cfg *config.Config) *ratelimiter.RateLimiter {

	dbRedisDataLimiter := database.NewRedisDataLimiter(database.NewRedisClient(cfg))
	rateLimiter := ratelimiter.NewLimiter(dbRedisDataLimiter, cfg.TokenMaxRequestsPerSecond, int64(cfg.LockDurationSeconds), int64(cfg.BlockDurationSeconds), int64(cfg.IPMaxRequestsPerSecond))

	if err := rateLimiter.RegisterPersonalizedTokens(context.Background()); err != nil {
		log.Fatal("Erro ao registrar o token:", err)
	}

	return rateLimiter
}
