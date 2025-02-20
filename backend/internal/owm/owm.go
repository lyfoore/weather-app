package owm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	weatherURL = "https://api.openweathermap.org/data/2.5/weather"
)

type owmResponse struct {
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
	Timezone int32  `json:"timezone"`
	Name     string `json:"name"`
	Weather  []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

type Response struct {
	Temp        float32 `json:"temp"`
	Time        string  `json:"timezone"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
}

func GetResponse(cityName string, APIkey *string) (*owmResponse, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", weatherURL, cityName, *APIkey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	owmResp := &owmResponse{}

	if err := json.NewDecoder(resp.Body).Decode(owmResp); err != nil {
		return nil, fmt.Errorf("JSON decode error: %w", err)
	}

	return owmResp, nil
}
