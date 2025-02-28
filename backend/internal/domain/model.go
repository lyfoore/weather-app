package domain

type ExternalAPIResponse struct {
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

type WeatherResponse struct {
	Temperature float32 `json:"temperature"`
	Timezone    int64   `json:"timezone"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Icon        string  `json:"icon"`
}
