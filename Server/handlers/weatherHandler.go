package handlers // Or your chosen package name

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

// --- Handler Function ---

// GetWeatherByCity handles requests to fetch specific weather data (City, Temp, Condition).
// It expects a query parameter 'city'.
func GetWeatherByCity(c *gin.Context) {
	// 1. Get the city name (or part of it) from the query parameter ?city=...
	queryCity := c.Query("city")
	if queryCity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'city' query parameter"})
		return
	}
	// Prepare for case-insensitive comparison
	queryCityLower := strings.ToLower(strings.TrimSpace(queryCity))

	// 2. Fetch the XML data from DHMZ
	resp, err := http.Get(dhmzURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data from DHMZ", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Check if DHMZ request was successful
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("DHMZ server returned non-OK status: %s", resp.Status)})
		return
	}

	// 3. Decode the XML into our structs
	var dhmzData Hrvatska
	decoder := xml.NewDecoder(resp.Body)
	// Optional: Handle potential non-UTF8 encodings if necessary
	// decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&dhmzData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse DHMZ XML data", "details": err.Error()})
		return
	}

	// 4. Filter results and build the response slice using CityWeatherInfo
	var results []CityWeatherInfo // Use the new struct for the results
	for _, grad := range dhmzData.Gradovi {
		// Case-insensitive check if GradIme contains the query string
		if strings.Contains(strings.ToLower(grad.GradIme), queryCityLower) {

			// Parse the temperature string to a float
			tempStr := strings.TrimSpace(grad.Podatci.Temp)
			tempFloat, err := strconv.ParseFloat(tempStr, 64)
			if err != nil {
				// Log the error but skip this city if temp is unparseable
				fmt.Printf("Warning: Could not parse temperature '%s' for city '%s': %v\n", tempStr, grad.GradIme, err)
				continue // Skip this city
			}

			// Create the result struct with only the requested fields
			results = append(results, CityWeatherInfo{
				City:             grad.GradIme,                            // Get City Name
				Temperature:      tempFloat,                               // Get parsed Temperature
				WeatherCondition: strings.TrimSpace(grad.Podatci.Vrijeme), // Get Weather Condition
			})
		}
	}

	// 5. Return the found results (can be an empty list if no match) as JSON
	c.JSON(http.StatusOK, results)
}
