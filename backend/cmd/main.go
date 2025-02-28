package main

import (
	"github.com/lyfoore/weather-app/config"
	"github.com/lyfoore/weather-app/internal/client"
	"github.com/lyfoore/weather-app/internal/service"
	transport "github.com/lyfoore/weather-app/internal/transport/http"
)

func main() {
	config := config.NewConfig()

	client := client.NewWeatherClient(config)
	service := service.NewWeatherService(client)
	router := transport.SetupRouter(service)

	router.Run()
}
