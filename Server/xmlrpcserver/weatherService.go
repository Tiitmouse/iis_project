package xmlrpcserver

import (
	"errors"
	"fmt"
	"iis_server/handlers"
	"net/http"
)

type WeatherService struct{}

type GetTemperatureArgs struct {
	City string
}

type GetTemperatureReply struct {
	Data []handlers.CityWeatherInfo
}

func (ws *WeatherService) GetTemperature(r *http.Request, args *GetTemperatureArgs, reply *GetTemperatureReply) error {
	city := args.City
	fmt.Printf("XML-RPC (gorilla) call received for city: %s\n", city)

	if city == "" {
		fmt.Println("Error: City parameter was empty after XML-RPC mapping.")
		return errors.New("city parameter is missing or empty")
	}

	dhmzData, err := handlers.FetchAndParseDHMZData()
	if err != nil {
		fmt.Printf("Error fetching/parsing DHMZ data for XML-RPC: %v\n", err)
		return fmt.Errorf("failed to retrieve weather data source: %w", err)
	}

	rez, err := handlers.FindTemperatureForCity(dhmzData, city)
	if err != nil {
		fmt.Printf("Error finding temperature for city '%s' via XML-RPC: %v\n", city, err)
		if errors.Is(err, handlers.ErrCityNotFound) {
			return fmt.Errorf("city '%s' not found", city)
		}
		return fmt.Errorf("error processing temperature for city '%s': %w", city, err)
	}

	fmt.Printf("Found temperature for %s: %v\n", city, rez)
	reply.Data = rez
	return nil
}
