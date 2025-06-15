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
	City             string  `json:"city"`
	Temperature      float64 `json:"temperature"`
	WeatherCondition string  `json:"weatherCondition"`
}

type XMLRPCMethodResponse struct {
	XMLName xml.Name      `xml:"methodResponse"`
	Params  []XMLRPCParam `xml:"params>param"`
	Fault   *XMLRPCFault  `xml:"fault,omitempty"`
}

type XMLRPCParam struct {
	Value XMLRPCValue `xml:"value"`
}

// XMLRPCValue updated to capture direct scalar types as well as structs
type XMLRPCValue struct {
	Struct []XMLRPCMember `xml:"struct>member,omitempty"`
	String string         `xml:"string,omitempty"`
	Double float64        `xml:"double,omitempty"` // This will capture <value><double>VALUE</double></value>
	Int    int            `xml:"int,omitempty"`    // For <int>
	I4     int            `xml:"i4,omitempty"`     // For <i4> (another common int type in XML-RPC)
}

type XMLRPCMember struct {
	Name  string         `xml:"name"`
	Value InnerDataValue `xml:"value"` // Value within a struct member
}

type InnerDataValue struct {
	String string  `xml:"string"`
	Double float64 `xml:"double"`
	Int    int     `xml:"int"`
	I4     int     `xml:"i4,omitempty"`
}

type XMLRPCFault struct {
	Value XMLRPCFaultValue `xml:"value"`
}

type XMLRPCFaultValue struct {
	Struct []XMLRPCMember `xml:"struct>member"`
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

	// The XML request body sends the city as a simple string parameter.
	// The server will map this positionally to the correct field in its argument struct.
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

	var rpcResp XMLRPCMethodResponse
	if err := xml.Unmarshal(body, &rpcResp); err != nil {
		log.Printf("FetchWeather (client-side): Error unmarshaling XML-RPC response: %v. Response body: %s", err, string(body))
		return nil, fmt.Errorf("error unmarshaling XML-RPC response: %w. Response body: %s", err, string(body))
	}

	if rpcResp.Fault != nil {
		faultCode := -1
		faultString := "Unknown XML-RPC fault"
		for _, member := range rpcResp.Fault.Value.Struct {
			if member.Name == "faultCode" {
				if member.Value.Int != 0 { // Prefer <int>
					faultCode = member.Value.Int
				} else { // Fallback for <i4> if server uses that for faultCode
					faultCode = member.Value.I4
				}
			} else if member.Name == "faultString" {
				faultString = member.Value.String
			}
		}
		log.Printf("FetchWeather (client-side): XML-RPC fault received: code %d, message: \"%s\"", faultCode, faultString)
		return nil, fmt.Errorf("XML-RPC fault: code %d, message: \"%s\"", faultCode, faultString)
	}

	var temperature float64
	foundTemperature := false

	if len(rpcResp.Params) > 0 {
		temperature = rpcResp.Params[0].Value.Double
		foundTemperature = true
		log.Printf("FetchWeather (client-side): Extracted temperature %.6f directly from response param.", temperature)
	}

	if !foundTemperature {
		log.Printf("FetchWeather (client-side): Temperature not found in XML-RPC response despite no fault. Response body: %s", string(body))
		return nil, fmt.Errorf("temperature not found in XML-RPC response. Response body: %s", string(body))
	}

	weatherInfo := CityWeatherInfo{
		City:        city,
		Temperature: temperature,
	}
	log.Printf("FetchWeather (client-side): Successfully processed weather for city '%s', temp: %.1f", city, temperature)

	return []CityWeatherInfo{weatherInfo}, nil
}
