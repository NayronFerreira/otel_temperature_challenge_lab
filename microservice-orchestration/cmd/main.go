package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/NayronFerreira/microservice-orchestration/config"
	"github.com/NayronFerreira/microservice-orchestration/internal/infra/web/controller"
	"github.com/NayronFerreira/microservice-orchestration/internal/infra/web/handlers"
	"github.com/NayronFerreira/microservice-orchestration/internal/infra/web/server"
	"github.com/NayronFerreira/microservice-orchestration/internal/service"
	"github.com/spf13/viper"
)

func main() {

	if err := config.LoadConfig("."); err != nil {
		log.Fatal(err)
	}

	signChannel := make(chan os.Signal, 1)
	signal.Notify(signChannel, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := config.InitializeOpenTelemetry()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown tracer provider: ", err)
		}
	}()

	server := server.NewWebServer(viper.GetString("WEB_PORT_SERVICE_ORCHESTRATION"))
	server.MountMiddlewares()

	handler := handlers.NewHandler(service.NewViaCepService(), service.NewWeatherApiService())
	controller := controller.NewController(server.Router, handler)
	controller.Route()

	go func() {
		server.Start()
	}()

	select {
	case <-signChannel:
		log.Println("shutting down server gracefully...")
	case <-ctx.Done():
		log.Println("shutting down server...")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}
