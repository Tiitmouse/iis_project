package main

import (
	"Desktop/api"
	"context"
	"fmt"

	"github.com/fiorix/wsdl2go/soap"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type JsFile struct {
	Name string
	Size int
	Type string
}

type Response struct {
	Data  any
	Error any
}

func (a *App) Validate(file JsFile, data string, m string) Response {
	fmt.Printf("file: %+v\n", file)
	rez, err := api.Validate(file.Name, []byte(data), m)
	if err != nil {
		return Response{Data: nil, Error: err.Error()}
	}
	return Response{Data: rez}
}

func (a *App) FetchWeather(city string) ([]api.CityWeatherInfo, error) {
	return api.FetchWeather(city)
}

func (a *App) SearchContacts(domain string) ([]api.ArrayOfContactRecord, error) {
	cli := soap.Client{
		URL:       "http://server",
		Namespace: api.Namespace,
	}
	soapService := api.NewIContactSearchService(&cli)
	contactReply, err := soapService.SearchContacts(&api.SearchContacts{
		SearchTerm: &domain,
	})
	if err != nil {
		return nil, err
	}
	return []api.ArrayOfContactRecord{*contactReply.SearchContactsResult}, nil
}
