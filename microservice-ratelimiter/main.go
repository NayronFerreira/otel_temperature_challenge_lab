package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/NayronFerreira/microservice-ratelimiter/config"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/database"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/handler"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/middleware"
	"github.com/NayronFerreira/microservice-ratelimiter/infra/web/server"
	"github.com/NayronFerreira/microservice-ratelimiter/ratelimiter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	InitializeOpenTelemetry(cfg)

	rateLimitedHandler := middleware.RateLimitMiddleware(handler.DummyHandler(), SetupRateLimiter(cfg), cfg)

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

func InitializeOpenTelemetry(cfg *config.Config) (func(ctx context.Context) error, error) {
	ctx := context.Background()

	serviceResource, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfg.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collectorConnection, err := grpc.DialContext(
		ctx,
		cfg.OtelCollectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Create a new trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(collectorConnection))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(traceExporter)
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(serviceResource),
		sdktrace.WithSpanProcessor(batchSpanProcessor),
	)

	// Set the global trace provider
	otel.SetTracerProvider(traceProvider)

	// Set the global text map propagator
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return traceProvider.Shutdown, nil
}
