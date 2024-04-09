package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/NayronFerreira/otel_temperature_challenge_lab/internal/web"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initProvider(ctx, config.ServiceName, config.CollectorURL)
	if err != nil {
		log.Fatalf("Failed to initialize OpenTelemetry: %v", err)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown OpenTelemetry: %v", err)
		}
	}()

	web.InitializeRoutes(config)
}

func initProvider(ctx context.Context, serviceName, urlCollector string) (func(context.Context) error, error) {
	resource, err := createResource(ctx, serviceName)
	if err != nil {
		return nil, err
	}

	conn, err := createGRPCConn(ctx, urlCollector)
	if err != nil {
		return nil, err
	}

	traceExporter, err := createTraceExporter(ctx, conn)
	if err != nil {
		return nil, err
	}

	traceProvider := createTraceProvider(resource, traceExporter)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return traceProvider.Shutdown, nil
}

func createResource(ctx context.Context, serviceName string) (*resource.Resource, error) {
	return resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
}

func createGRPCConn(ctx context.Context, urlCollector string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return grpc.DialContext(
		ctx,
		urlCollector,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
}

func createTraceExporter(ctx context.Context, conn *grpc.ClientConn) (*otlptracegrpc.Exporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func createTraceProvider(resource *resource.Resource, traceExporter *otlptracegrpc.Exporter) *sdktrace.TracerProvider {
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource),
		sdktrace.WithSpanProcessor(bsp),
	)
}
ß