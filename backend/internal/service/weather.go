package service

import (
	"log"

	"github.com/lyfoore/weather-app/internal/client"
	"github.com/lyfoore/weather-app/internal/domain"
)

type WeatherService struct {
	client client.ExternalWeatherClient
}

func NewWeatherService(c *client.ExternalWeatherClient) *WeatherService {
	return &WeatherService{
		client: *c,
	}
}

func (s WeatherService) ProcessWeatherResponse(cityName string) (*domain.WeatherResponse, error) {
	externalResponse, err := s.client.GetExternalResponse(cityName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &domain.WeatherResponse{
		Temperature: externalResponse.Main.Temp,
		Timezone:    externalResponse.Timezone,
		Description: externalResponse.Weather[0].Description,
		Name:        externalResponse.Name,
		Icon:        externalResponse.Weather[0].Icon,
	}, nil
}
