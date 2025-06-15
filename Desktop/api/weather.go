package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const xmlRPCWeatherURL = "http://localhost:8089/RPC2"

type CityWeatherInfo struct {
	City             string
	Temperature      float64
	WeatherCondition string
}

type MethodResponse struct {
	XMLName xml.Name `xml:"methodResponse"`
	Cities  []struct {
		Members []Member `xml:"member"`
	} `xml:"params>param>value>array>data>value>struct"`
}

type Member struct {
	Name  string `xml:"name"`
	Value struct {
		String string  `xml:"string"`
		Double float64 `xml:"double"`
	} `xml:"value"`
}

func FetchWeather(city string) ([]CityWeatherInfo, error) {
	log.Printf("FetchWeather (client-side): Received city parameter: '%s'", city)

	var escapedCityBuffer bytes.Buffer
	if err := xml.EscapeText(&escapedCityBuffer, []byte(city)); err != nil {
		log.Printf("FetchWeather (client-side): Error escaping city name '%s': %v", city, err)
		return nil, fmt.Errorf("error escaping city name for XML: %w", err)
	}
	escapedCity := escapedCityBuffer.String()
	log.Printf("FetchWeather (client-side): Escaped city name: '%s'", escapedCity)

	xmlRequestBody := fmt.Sprintf(`<?xml version="1.0"?>
<methodCall>
  <methodName>WeatherService.GetTemperature</methodName>
  <params>
    <param>
      <value><string>%s</string></value>
    </param>
  </params>
</methodCall>`, escapedCity)
	log.Printf("FetchWeather (client-side): Constructed XML-RPC request body:\n%s", xmlRequestBody)

	log.Printf("FetchWeather (client-side): Attempting XML-RPC call to URL: %s", xmlRPCWeatherURL)
	reqBody := strings.NewReader(xmlRequestBody)
	resp, err := http.Post(xmlRPCWeatherURL, "text/xml; charset=utf-8", reqBody)
	if err != nil {
		log.Printf("FetchWeather (client-side): http.Post error to %s: %v", xmlRPCWeatherURL, err)
		return nil, fmt.Errorf("error making XML-RPC request to %s: %w", xmlRPCWeatherURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("FetchWeather (client-side): Error reading XML-RPC response body: %v", err)
		return nil, fmt.Errorf("error reading XML-RPC response body: %w", err)
	}
	log.Printf("FetchWeather (client-side): Received raw XML-RPC response body:\n%s", string(body))

	if resp.StatusCode != http.StatusOK {
		log.Printf("FetchWeather (client-side): XML-RPC request to %s failed with status %d. Response body: %s", xmlRPCWeatherURL, resp.StatusCode, string(body))
		return nil, fmt.Errorf("XML-RPC request failed with HTTP status %d: %s", resp.StatusCode, string(body))
	}

	var rpcResp MethodResponse
	if err := xml.Unmarshal(body, &rpcResp); err != nil {
		log.Printf("FetchWeather (client-side): Error unmarshaling XML-RPC response: %v. Response body: %s", err, string(body))
		return nil, fmt.Errorf("error unmarshaling XML-RPC response: %w. Response body: %s", err, string(body))
	}
	data := transformToCityWeatherInfo(rpcResp)
	fmt.Printf("rpcResp: %+v\n", data)
	return data, nil
}

func transformToCityWeatherInfo(resp MethodResponse) []CityWeatherInfo {
	// Initialize an empty slice to hold the final, clean data.
	weatherInfos := make([]CityWeatherInfo, 0, len(resp.Cities))

	// Iterate over each city struct that was unmarshaled from the XML.
	for _, cityStruct := range resp.Cities {
		var info CityWeatherInfo
		// For each city, iterate over its members (City, Temperature, WeatherCondition).
		for _, member := range cityStruct.Members {
			// Use a switch statement to assign the correct value to the corresponding field in the info struct.
			switch member.Name {
			case "City":
				info.City = member.Value.String
			case "Temperature":
				info.Temperature = member.Value.Double
			case "WeatherCondition":
				info.WeatherCondition = member.Value.String
			}
		}
		// Append the fully populated, clean struct to the results slice.
		weatherInfos = append(weatherInfos, info)
	}

	return weatherInfos
}
