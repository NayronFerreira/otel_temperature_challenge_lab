package main

import (
	"fmt"

	"github.com/NayronFerreira/otel_temperature_challenge_lab/config"
	"github.com/NayronFerreira/otel_temperature_challenge_lab/web"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	web.InitializeRoutes(config)

}
