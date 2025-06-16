package main

import (
	"Desktop/api"
	"context"
	"fmt"
	"os"
	"os/exec" // Added for running external commands
	"path/filepath"

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
	fmt.Printf("ManualSearch called with domain: %s\n", domain)
	result, err := api.ManualRequestSearchContacts(domain)
	if err != nil {
		fmt.Printf("ManualSearch error: %v\n", err)
		return nil, err
	}
	fmt.Printf("ManualSearch result: %+v\n", result)
	return result, nil
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

func (a *App) CheckSoapFileExists() bool {

	filePath := "../SoapService/WebsiteContactsService/contacts_data.xml"

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("Error getting absolute path for %s: %v\\n", filePath, err)
		return false
	}

	fmt.Printf("Checking for file at absolute path: %s\\n", absPath)
	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		fmt.Printf("File not found: %s\\n", absPath)
		return false
	}
	if err != nil {
		fmt.Printf("Error checking file %s: %v\\n", absPath, err)
		return false
	}
	fmt.Printf("File found: %s\\n", absPath)
	return true
}

func (a *App) RunJaxbValidation() (string, error) {
	fmt.Println("RunJaxbValidation called")

	jarPath, err := filepath.Abs("../JaxbThing/dist/JaxbThing.jar")
	if err != nil {
		fmt.Printf("Error getting absolute path for JaxbThing.jar: %v\\n", err)
		return "", fmt.Errorf("error resolving JaxbThing.jar path: %w", err)
	}
	fmt.Printf("Attempting to run JAR: %s\\n", jarPath)

	if _, err := os.Stat(jarPath); os.IsNotExist(err) {
		errMsg := fmt.Sprintf("JaxbThing.jar not found at %s. Please ensure the Java project is built and the JAR is in the 'dist' directory.", jarPath)
		fmt.Println(errMsg)
		return "", fmt.Errorf("%s", errMsg)
	}

	workingDir, err := filepath.Abs("../JaxbThing")
	if err != nil {
		fmt.Printf("Error getting absolute path for JaxbThing working directory: %v\\n", err)
		return "", fmt.Errorf("error resolving JaxbThing working directory: %w", err)
	}
	fmt.Printf("Setting working directory for Java app to: %s\\n", workingDir)

	// java -cp JaxbThing.jar jaxbThing.JaxbThing
	cmd := exec.Command("java", "-jar", jarPath)
	cmd.Dir = workingDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running JAXB validation: %v\\nOutput: %s\\n", err, string(output))
		return "", fmt.Errorf("failed to run JAXB validation: %w. Output: %s", err, string(output))
	}

	fmt.Printf("JAXB validation output: %s\\n", string(output))
	return string(output), nil
}
