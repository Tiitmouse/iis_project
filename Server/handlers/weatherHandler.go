package handlers

import (
	"encoding/xml"
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

func GetWeatherByCity(c *gin.Context) {
	queryCity := c.Query("city")
	if queryCity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'city' query parameter"})
		return
	}
	queryCityLower := strings.ToLower(strings.TrimSpace(queryCity))

	resp, err := http.Get(dhmzURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from DHMZ", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("DHMZ server returned non-OK status: %s", resp.Status)})
		return
	}

	var dhmzData Hrvatska
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&dhmzData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse DHMZ XML data", "details": err.Error()})
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

	c.JSON(http.StatusOK, results)
}
