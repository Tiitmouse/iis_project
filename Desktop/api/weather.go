package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const weatherURL = "http://localhost:8088/weather?city=%s"

type CityWeatherInfo struct {
	City             string  `json:"city"`
	Temperature      float64 `json:"temperature"`
	WeatherCondition string  `json:"weatherCondition"`
}

func FetchWeather(city string) ([]CityWeatherInfo, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf(weatherURL, encodedCity)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var weatherData []CityWeatherInfo
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return weatherData, nil
}
