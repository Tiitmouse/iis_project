package validator

import (
	"errors"
	"fmt"

	"github.com/killi1812/libxml2"
	"github.com/killi1812/libxml2/relaxng"
	"github.com/killi1812/libxml2/types"
	"github.com/killi1812/libxml2/xsd"
)

var ErrBadSyntax = errors.New("failed to parse XML, syntax error")

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
		return errors.Join(ErrBadSyntax, err)
	}
	defer doc.Free()

	if err := schema.Validate(doc); err != nil {
		var validationError types.SchemaValidationError
		if errors.As(err, &validationError) {
			return fmt.Errorf("schema not valid: %v", validationError.Errors)
		}
		return fmt.Errorf("schema validation failed: %w", err)
	}

	return nil
}
