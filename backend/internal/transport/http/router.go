package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/lyfoore/weather-app/internal/service"
)

type WeatherServer struct{}

func SetupRouter(weatherService *service.WeatherService) *gin.Engine {
	router := gin.Default()

	weatherHandler := NewWeatherHandler(weatherService)

	api := router.Group("/api")
	{
		api.GET("/getWeather/:cityName", weatherHandler.GetWeather)
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	return router
}
