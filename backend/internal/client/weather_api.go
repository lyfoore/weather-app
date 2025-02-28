package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lyfoore/weather-app/config"
	"github.com/lyfoore/weather-app/internal/domain"
)

const (
	weatherURL = "https://api.openweathermap.org/data/2.5/weather"
)

type ExternalWeatherClient struct {
	cfg config.Config
}

func NewWeatherClient(c *config.Config) *ExternalWeatherClient {
	return &ExternalWeatherClient{
		cfg: *c,
	}
}

func (client *ExternalWeatherClient) GetExternalResponse(cityName string) (*domain.ExternalAPIResponse, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", weatherURL, cityName, client.cfg.APIkey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("RESPONSE:", cityName)

	var data domain.ExternalAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln("Error while decoding externalAPIResponse", err)
		return nil, err
	}

	return &data, nil
}
