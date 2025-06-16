package storage

import (
	"embed"
	"encoding/xml"
	"fmt"
	"iis_server/structs"

	"github.com/google/uuid"
)

//go:embed data.xml
var seedFS embed.FS

type XMLRoot struct {
	XMLName   xml.Name `xml:"root"`
	Status    string   `xml:"status"`
	RequestID string   `xml:"request_id"`
	Data      XMLData  `xml:"data"`
}
type XMLData struct {
	XMLName      xml.Name         `xml:"data"`
	Domain       string           `xml:"domain"`
	Query        string           `xml:"query"`
	Emails       []XMLEmail       `xml:"emails"`
	PhoneNumbers []XMLPhoneNumber `xml:"phone_numbers"`
	Facebook     string           `xml:"facebook,omitempty"`
	Instagram    string           `xml:"instagram,omitempty"`
	Tiktok       string           `xml:"tiktok,omitempty"`
	Snapchat     string           `xml:"snapchat,omitempty"`
	Twitter      string           `xml:"twitter,omitempty"`
	Linkedin     string           `xml:"linkedin,omitempty"`
	Github       string           `xml:"github,omitempty"`
	Youtube      string           `xml:"youtube,omitempty"`
	Pinterest    string           `xml:"pinterest,omitempty"`
}
type XMLEmail struct {
	XMLName xml.Name `xml:"emails"`
	Value   string   `xml:"value"`
	Sources []string `xml:"sources"`
}
type XMLPhoneNumber struct {
	XMLName xml.Name `xml:"phone_numbers"`
	Value   string   `xml:"value"`
	Sources []string `xml:"sources"`
}

func parseAndTransformSeedXML() ([]structs.Contact, error) {
	seedFileName := "data.xml"
	xmlBytes, err := seedFS.ReadFile(seedFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded seed file %s: %w", seedFileName, err)
	}

	var parsedData XMLRoot
	err = xml.Unmarshal(xmlBytes, &parsedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal seed XML from %s: %w", seedFileName, err)
	}

	var contacts []structs.Contact

	for _, email := range parsedData.Data.Emails {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.EmailType, Value: email.Value, Sources: email.Sources, Name: ""})
	}
	for _, phone := range parsedData.Data.PhoneNumbers {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.PhoneType, Value: phone.Value, Sources: phone.Sources, Name: ""})
	}
	if parsedData.Data.Facebook != "" {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.SocialType, Value: parsedData.Data.Facebook, Name: "facebook", Sources: []string{}})
	}
	if parsedData.Data.Instagram != "" {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.SocialType, Value: parsedData.Data.Instagram, Name: "instagram", Sources: []string{}})
	}
	if parsedData.Data.Twitter != "" {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.SocialType, Value: parsedData.Data.Twitter, Name: "twitter", Sources: []string{}})
	}
	if parsedData.Data.Linkedin != "" {
		contacts = append(contacts, structs.Contact{ID: uuid.NewString(), Type: structs.SocialType, Value: parsedData.Data.Linkedin, Name: "linkedin", Sources: []string{}})
	}

	return contacts, nil
}

func SeedStore(store *ContactStore) {
	fmt.Println("INFO: Parsing and seeding contact store from embedded XML file...")
	seeds, err := parseAndTransformSeedXML()
	if err != nil {
		fmt.Printf("ERROR: Failed to parse seed data from embedded XML: %v\n", err)
		fmt.Println("WARN: Starting with an empty contact store.")
		return
	}

	addedCount := 0
	for _, contact := range seeds {
		_, addErr := store.AddContact(contact)
		if addErr != nil {
			fmt.Printf("ERROR: Failed to add seed contact ID %s: %v\n", contact.ID, addErr)
		} else {
			addedCount++
		}
	}
	fmt.Printf("INFO: Embedded XML Seeding complete. Added: %d contacts.\n", addedCount)
}
