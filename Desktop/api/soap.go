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
const dataContractNamespace = "http://schemas.datacontract.org/2004/07/WebsiteContactsService.Models"
const arraysNamespace = "http://schemas.microsoft.com/2003/10/Serialization/Arrays"

type SoapEmailEntry struct {
	XMLName xml.Name `xml:",omitempty"`
	Value   string   `xml:"Value"`
	Sources []string `xml:"Sources>string"`
}

type SoapPhoneEntry struct {
	XMLName xml.Name `xml:",omitempty"`
	Value   string   `xml:"Value"`
	Sources []string `xml:"Sources>string"`
}

type SoapContactRecord struct {
	XMLName      xml.Name         `xml:"Contact"`
	Domain       string           `xml:"Domain"`
	Query        string           `xml:"Query"`
	Emails       []SoapEmailEntry `xml:"Emails>EmailEntry"`
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

type SoapSearchContactsResult struct {
	XMLName xml.Name            `xml:"SearchContactsResult"`
	Contact []SoapContactRecord `xml:"Contact"`
}

type SoapSearchResponse struct {
	XMLName              xml.Name                 `xml:"SearchContactsResponse"`
	SearchContactsResult SoapSearchContactsResult `xml:"SearchContactsResult"`
}

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
	log.Printf("SOAP Response Body: %s", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SOAP request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var envelope struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			XMLName                xml.Name           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
			SearchContactsResponse SoapSearchResponse `xml:"http://tempuri.org/ SearchContactsResponse"`
		} `xml:"Body"`
	}

	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Printf("SOAP Client: Error unmarshaling response: %v. Raw Response: %s", err, string(body))
		return nil, fmt.Errorf("error parsing SOAP response XML: %w", err)
	}
	log.Printf("Parsed envelope.Body.SearchContactsResponse: %+v", envelope.Body.SearchContactsResponse)

	return envelope.Body.SearchContactsResponse.SearchContactsResult.Contact, nil
}
