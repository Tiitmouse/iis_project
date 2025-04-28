package handlers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Hrvatska struct {
	XMLName xml.Name `xml:"Hrvatska"`
	Gradovi []Grad   `xml:"Grad"`
}

type Grad struct {
	XMLName xml.Name `xml:"Grad"`
	GradIme string   `xml:"GradIme"`
	Podatci Podatci  `xml:"Podatci"`
}

type Podatci struct {
	XMLName xml.Name `xml:"Podatci"`
	Temp    string   `xml:"Temp"`
	Vrijeme string   `xml:"Vrijeme"`
}

type CityWeatherInfo struct {
	City             string  `json:"city"`
	Temperature      float64 `json:"temperature"`
	WeatherCondition string  `json:"weatherCondition"`
}

const dhmzURL = "https://vrijeme.hr/hrvatska_n.xml"

var ErrCityNotFound = errors.New("city not found in DHMZ data")

func FetchAndParseDHMZData() (*Hrvatska, error) {
	resp, err := http.Get(dhmzURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from DHMZ: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DHMZ server returned non-OK status: %s", resp.Status)
	}

	var dhmzData Hrvatska
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&dhmzData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DHMZ XML data: %w", err)
	}
	return &dhmzData, nil
}

func FindTemperatureForCity(data *Hrvatska, cityName string) (float64, error) {
	queryCityLower := strings.ToLower(strings.TrimSpace(cityName))
	for _, grad := range data.Gradovi {
		if strings.EqualFold(strings.ToLower(grad.GradIme), queryCityLower) {
			tempStr := strings.TrimSpace(grad.Podatci.Temp)
			tempFloat, err := strconv.ParseFloat(tempStr, 64)
			if err != nil {
				return 0, fmt.Errorf("could not parse temperature '%s' for city '%s': %w", tempStr, grad.GradIme, err)
			}
			return tempFloat, nil
		}
	}
	return 0, ErrCityNotFound
}

func GetWeatherByCity(c *gin.Context) {
	queryCity := c.Query("city")
	if queryCity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'city' query parameter"})
		return
	}
	queryCityLower := strings.ToLower(strings.TrimSpace(queryCity))

	dhmzData, err := FetchAndParseDHMZData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get DHMZ data", "details": err.Error()})
		return
	}

	var results []CityWeatherInfo
	for _, grad := range dhmzData.Gradovi {
		if strings.Contains(strings.ToLower(grad.GradIme), queryCityLower) {
			tempStr := strings.TrimSpace(grad.Podatci.Temp)
			tempFloat, err := strconv.ParseFloat(tempStr, 64)
			if err != nil {
				fmt.Printf("Warning: Could not parse temperature '%s' for city '%s': %v\n", tempStr, grad.GradIme, err)
				continue
			}
			results = append(results, CityWeatherInfo{
				City:             grad.GradIme,
				Temperature:      tempFloat,
				WeatherCondition: strings.TrimSpace(grad.Podatci.Vrijeme),
			})
		}
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No weather data found for city containing '%s'", queryCity)})
		return
	}

	c.JSON(http.StatusOK, results)
}
