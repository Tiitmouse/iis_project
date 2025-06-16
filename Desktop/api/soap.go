package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const manualSoapServiceURL = "http://localhost:5157/contact"

// Define the namespace that the .NET SOAP service uses for its data contracts.
const dataContractNamespace = "http://schemas.datacontract.org/2004/07/WebsiteContactsService.Models"
const arraysNamespace = "http://schemas.microsoft.com/2003/10/Serialization/Arrays"

// SoapEmailEntry matches the <d4p1:EmailEntry> element.
type SoapEmailEntry struct {
	XMLName xml.Name `xml:",omitempty"` // The actual name is specified in SoapContactRecord's tag.
	Value   string   `xml:"Value"`      // This will inherit namespace from parent <EmailEntry> if not specified.
	// Correctly parse <d8p1:string> elements within <d4p1:Sources>
	Sources []string `xml:"Sources>string"`
}

// SoapPhoneEntry matches the <d4p1:PhoneEntry> element.
type SoapPhoneEntry struct {
	XMLName xml.Name `xml:",omitempty"`
	Value   string   `xml:"Value"`
	Sources []string `xml:"Sources>string"`
}

// SoapContactRecord mirrors the C# Contact class and <d4p1:Contact> element.
type SoapContactRecord struct {
	XMLName xml.Name `xml:"Contact"`
	Domain  string   `xml:"Domain"`
	Query   string   `xml:"Query"`
	// Collection <d4p1:Emails> contains repeating <d4p1:EmailEntry> items.
	Emails []SoapEmailEntry `xml:"Emails>EmailEntry"`
	// Collection <d4p1:PhoneNumbers> contains repeating <d4p1:PhoneEntry> items.
	PhoneNumbers []SoapPhoneEntry `xml:"PhoneNumbers>PhoneEntry"`
	Facebook     string           `xml:"Facebook,omitempty"`
	Instagram    string           `xml:"Instagram,omitempty"`
	Github       string           `xml:"Github,omitempty"`
	Linkedin     string           `xml:"Linkedin,omitempty"`
	Twitter      string           `xml:"Twitter,omitempty"`
	Youtube      string           `xml:"Youtube,omitempty"`
	Pinterest    string           `xml:"Pinterest,omitempty"`
	Tiktok       string           `xml:"Tiktok,omitempty"`
	Snapchat     string           `xml:"Snapchat,omitempty"`
}

// SoapSearchContactsResult contains the list of <d4p1:Contact> records.
// This element <SearchContactsResult> is in the "http://tempuri.org/" namespace.
type SoapSearchContactsResult struct {
	XMLName xml.Name `xml:"SearchContactsResult"`
	// Each <Contact> item is in the dataContractNamespace implicitly via its own struct tag or by Go's XML rules.
	Contact []SoapContactRecord `xml:"Contact"`
}

// SoapSearchResponse is the top-level structure for the useful part of the SOAP body.
// This element <SearchContactsResponse> is in the "http://tempuri.org/" namespace.
type SoapSearchResponse struct {
	XMLName              xml.Name                 `xml:"SearchContactsResponse"`
	SearchContactsResult SoapSearchContactsResult `xml:"SearchContactsResult"`
}

// ManualRequestSearchContacts builds and sends a manual SOAP request.
func ManualRequestSearchContacts(domain string) ([]SoapContactRecord, error) {
	log.Printf("Executing Manual SOAP Request for domain: '%s'", domain)

	requestBody := fmt.Sprintf(`
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
   <soapenv:Header/>
   <soapenv:Body>
      <tem:SearchContacts>
         <tem:searchTerm>%s</tem:searchTerm>
      </tem:SearchContacts>
   </soapenv:Body>
</soapenv:Envelope>`, domain)

	req, err := http.NewRequest("POST", manualSoapServiceURL, strings.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating SOAP request: %w", err)
	}

	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("SOAPAction", `"http://tempuri.org/IContactSearchService/SearchContacts"`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending SOAP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading SOAP response body: %w", err)
	}
	log.Printf("SOAP Response Body: %s", string(body)) // Crucial log!

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SOAP request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var envelope struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"` // s:Envelope
		Body    struct {
			XMLName                xml.Name           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"` // s:Body
			SearchContactsResponse SoapSearchResponse `xml:"http://tempuri.org/ SearchContactsResponse"`     // xmlns="http://tempuri.org/"
		} `xml:"Body"`
	}

	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Printf("SOAP Client: Error unmarshaling response: %v. Raw Response: %s", err, string(body))
		return nil, fmt.Errorf("error parsing SOAP response XML: %w", err)
	}
	log.Printf("Parsed envelope.Body.SearchContactsResponse: %+v", envelope.Body.SearchContactsResponse)

	// Return the list of contact records
	return envelope.Body.SearchContactsResponse.SearchContactsResult.Contact, nil
}
