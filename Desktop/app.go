package main

import (
	"Desktop/api"
	"context"
	"fmt"

	"github.com/fiorix/wsdl2go/soap"
)

type App struct {
	ctx     context.Context
	service api.IContactSearchService
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	cli := &soap.Client{
		URL:       "http://localhost:5157/contact",
		Namespace: api.Namespace,
	}
	a.service = api.NewIContactSearchService(cli)
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// func (a *App) SearchContacts(domain string) (*api.ArrayOfContactRecord, error) {
// 	fmt.Printf("(Unused) Searching for domain: %s\n", domain)
// 	contactReply, err := a.service.SearchContacts(&api.SearchContacts{
// 		SearchTerm: &domain,
// 	})
// 	if err != nil {
// 		fmt.Printf("SOAP call error: %v\n", err)
// 		return nil, err
// 	}
// 	return contactReply.SearchContactsResult, nil
// }

func (a *App) ManualSearch(domain string) ([]api.SoapContactRecord, error) {
	return api.ManualRequestSearchContacts(domain)
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
