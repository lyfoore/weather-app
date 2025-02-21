package owm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	weatherURL = "https://api.openweathermap.org/data/2.5/weather"
)

type owmResponse struct {
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
	Timezone int64  `json:"timezone"`
	Name     string `json:"name"`
	Weather  []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
}

type Response struct {
	Temperature float32 `json:"temperature"`
	Time        string  `json:"time"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
}

func GetResponse(cityName string, APIkey *string) (*Response, error) {
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

	response := &Response{}
	owmResp.convertToResponse(response)

	return response, nil
}

func (owmResp owmResponse) convertToResponse(resp *Response) {
	resp.Temperature = owmResp.Main.Temp
	resp.Description = owmResp.Weather[0].Description
	resp.Name = owmResp.Name
	resp.Icon = owmResp.Weather[0].Icon
	resp.Time = getCityTime(owmResp.Timezone)
}

func getCityTime(timeshift int64) string {
	currTime := time.Now().UTC().Add(time.Duration(timeshift) * time.Second)
	return currTime.Format(time.TimeOnly)
}
