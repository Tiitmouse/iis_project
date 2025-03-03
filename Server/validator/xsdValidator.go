package validator

import (
	"fmt"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func ValidateWithXSD(xmlData []byte, xsdFilePath string) error {
	schema, err := xsd.ParseFromFile(xsdFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse XSD file: %w", err)
	}
	defer schema.Free()

	doc, err := libxml2.Parse(xmlData)
	if err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}
	defer doc.Free()

	if err := schema.Validate(doc); err != nil {
		// TODO unwrap validators
		return fmt.Errorf("XML validation failed: %w", err)
	}

	return nil
}
