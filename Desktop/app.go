package main

import (
	"Desktop/api"
	"context"
	"fmt"

	"github.com/fiorix/wsdl2go/soap"
)

// App struct
type App struct {
	ctx     context.Context
	service api.IContactSearchService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	cli := &soap.Client{
		URL:       "http://localhost:5157/contact",
		Namespace: api.Namespace,
	}
	a.service = api.NewIContactSearchService(cli)

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

	fmt.Printf("searching %s\n", domain)
	contactReply, err := a.service.SearchContacts(&api.SearchContacts{
		SearchTerm: &domain,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return []api.ArrayOfContactRecord{*contactReply.SearchContactsResult}, nil
}
