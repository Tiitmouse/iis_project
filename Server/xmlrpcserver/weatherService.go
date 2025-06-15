package xmlrpcserver

import (
	"errors"
	"fmt"
	"iis_server/handlers"
	"net/http"
)

type WeatherService struct{}

// Define a struct for the arguments, which is what gorilla/rpc expects.
type GetTemperatureArgs struct {
	City string
}

type GetTemperatureReply struct {
	Temperature float64
}

// The method signature uses the *GetTemperatureArgs struct.
// The library will now automatically map the incoming string parameter from the client
// to the 'City' field of this struct.
func (ws *WeatherService) GetTemperature(r *http.Request, args *GetTemperatureArgs, reply *GetTemperatureReply) error {
	// Extract the city from the args struct
	city := args.City
	fmt.Printf("XML-RPC (gorilla) call received for city: %s\n", city)

	// Check if the city is empty, which would happen if mapping failed
	if city == "" {
		fmt.Println("Error: City parameter was empty after XML-RPC mapping.")
		return errors.New("city parameter is missing or empty")
	}

	dhmzData, err := handlers.FetchAndParseDHMZData()
	if err != nil {
		fmt.Printf("Error fetching/parsing DHMZ data for XML-RPC: %v\n", err)
		return fmt.Errorf("failed to retrieve weather data source: %w", err)
	}

	temperature, err := handlers.FindTemperatureForCity(dhmzData, city)
	if err != nil {
		fmt.Printf("Error finding temperature for city '%s' via XML-RPC: %v\n", city, err)
		if errors.Is(err, handlers.ErrCityNotFound) {
			return fmt.Errorf("city '%s' not found", city)
		}
		return fmt.Errorf("error processing temperature for city '%s': %w", city, err)
	}

	fmt.Printf("Found temperature for %s: %.1f\n", city, temperature)
	reply.Temperature = temperature
	return nil
}
