package validator

import (
	"fmt"

	"github.com/killi1812/libxml2"
	"github.com/killi1812/libxml2/relaxng"
	"github.com/killi1812/libxml2/types"
	"github.com/killi1812/libxml2/xsd"
)

func ValidateWithRNG(xmlData []byte, rngFilePath string) error {
	schema, err := relaxng.ParseFromFile(rngFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse RNG file: %w", err)
	}
	defer schema.Free()

	return validate(xmlData, schema)
}

func ValidateWithXSD(xmlData []byte, xsdFilePath string) error {
	schema, err := xsd.ParseFromFile(xsdFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse XSD file: %w", err)
	}
	defer schema.Free()

	return validate(xmlData, schema)
}

func validate(xmlData []byte, schema types.Schema) error {
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
