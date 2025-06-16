package main

import (
	"Desktop/api"
	"context"
	"fmt"
	"os" // Added for file operations

	// Added for running external commands
	"path/filepath" // Added for path manipulation

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

	filePath := "../SoapService/WebsiteContactsService/contacts_data.xml" // Adjusted path

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

func (a *App) RunJaxbValidation(xmlFilePath string) (string, error) {
	fmt.Printf("RunJaxbValidation called with XML file: %s\\n", xmlFilePath)

	// --- Placeholder for Java Application Execution ---
	// You will need to replace this with the actual command to run your Java application.
	//
	// Example considerations:
	// 1. Path to your Java executable (java).
	// 2. Path to your JAXB application's JAR file.
	// 3. Any command-line arguments your Java app needs (e.g., the xmlFilePath).
	//
	// Example (conceptual):
	// javaCmd := "java"
	// jarPath := "/path/to/your/JaxbThing.jar" // Replace with the actual path to your compiled JAR
	//
	// // Ensure xmlFilePath is an absolute path or accessible by the Java app
	// cmd := exec.Command(javaCmd, "-jar", jarPath, xmlFilePath)
	//
	// // Optional: Set working directory if your Java app expects to be run from a specific location
	// // cmd.Dir = "/path/to/your/JaxbThing/directory"
	//
	// output, err := cmd.CombinedOutput() // Gets both stdout and stderr
	// if err != nil {
	//   fmt.Printf("Error running JAXB validation: %v\\nOutput: %s\\n", err, string(output))
	//   return "", fmt.Errorf("failed to run JAXB validation: %w. Output: %s", err, string(output))
	// }
	//
	// fmt.Printf("JAXB validation output: %s\\n", string(output))
	// return string(output), nil
	// --- End Placeholder ---

	// For now, returning a dummy success message.
	// Replace this with the actual execution logic above.
	dummyOutput := fmt.Sprintf("Successfully processed (placeholder) %s. Implement Java call.", xmlFilePath)
	fmt.Println(dummyOutput)
	return dummyOutput, nil
}
