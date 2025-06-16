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

// SoapContactRecord is a renamed struct to avoid conflict with the one in search.go.
type SoapContactRecord struct {
	RecordType string   `xml:"RecordType"`
	Value      string   `xml:"Value"`
	Sources    []string `xml:"Sources>string"`
}

// SoapSearchResponse is a renamed struct for unmarshaling the SOAP response.
type SoapSearchResponse struct {
	XMLName       xml.Name            `xml:"SearchContactsResponse"`
	ContactRecord []SoapContactRecord `xml:"SearchContactsResult>ContactRecord"`
}

// ManualRequestSearchContacts builds and sends a manual SOAP request based on your curl command.
func ManualRequestSearchContacts(domain string) ([]SoapContactRecord, error) {
	log.Printf("Executing Manual SOAP Request for domain: '%s'", domain)

	// 1. Construct the SOAP XML payload exactly as in your curl command.
	requestBody := fmt.Sprintf(`
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
   <soapenv:Header/>
   <soapenv:Body>
      <tem:SearchContacts>
         <tem:searchTerm>%s</tem:searchTerm>
      </tem:SearchContacts>
   </soapenv:Body>
</soapenv:Envelope>`, domain)

	// 2. Create a new HTTP request.
	req, err := http.NewRequest("POST", manualSoapServiceURL, strings.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating SOAP request: %w", err)
	}

	// 3. Set headers exactly as in your curl command.
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("SOAPAction", `"http://tempuri.org/IContactSearchService/SearchContacts"`)

	// 4. Send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending SOAP request: %w", err)
	}
	defer resp.Body.Close()
	log.Printf("resp: %+v", resp)

	// 5. Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading SOAP response body: %w", err)
	}
	log.Printf("envelope: %+v", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SOAP request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// 6. Parse the XML response into our renamed structs.
	var envelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			Content SoapSearchResponse `xml:",any"`
		} `xml:"Body"`
	}

	if err := xml.Unmarshal(body, &envelope); err != nil {
		log.Printf("SOAP Client: Error unmarshaling response: %v. Raw Response: %s", err, string(body))
		return nil, fmt.Errorf("error parsing SOAP response XML: %w", err)
	}
	log.Printf("envelope: %+v", envelope)
	// 7. Return the clean list of contact records.
	return envelope.Body.Content.ContactRecord, nil
}
