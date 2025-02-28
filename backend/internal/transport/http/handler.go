package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyfoore/weather-app/internal/service"
)

type WeatherHandler struct {
	service service.WeatherService
}

func NewWeatherHandler(s *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		service: *s,
	}
}

func (h *WeatherHandler) GetWeather(ctx *gin.Context) {
	response, err := h.service.ProcessWeatherResponse(ctx.Param("cityName"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
