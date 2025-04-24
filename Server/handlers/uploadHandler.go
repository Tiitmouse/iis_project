package handlers

import (
	"errors"
	"fmt"
	"iis_server/validator"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/killi1812/libxml2/types"
)

func HandleXMLUpload(c *gin.Context) {
	xsdFilePath := "./schemas/getXSDschema.xsd"
	rngFilePath := "./schemas/getRNGschema.rng"

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get file err: %s", err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.String(http.StatusInternalServerError, "cannot open the file: %s", err.Error())
	}

	data := make([]byte, fileHeader.Size)
	_, err = file.Read(data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.String(http.StatusInternalServerError, "cannot read the file: %s", err.Error())
	}

	parts := strings.Split(c.Request.URL.Path, "/")
	validationType := parts[len(parts)-1]

	switch validationType {
	case "xsd":
		err = validator.ValidateWithXSD(data, xsdFilePath)

	case "rng":
		err = validator.ValidateWithRNG(data, rngFilePath)

	default:
		c.String(http.StatusInternalServerError, "Invalid validation: %s, %s", parts[len(parts)-1], err.Error())
	}

	if err != nil {
		var validationError types.SchemaValidationError
		if errors.As(err, &validationError) {
			c.String(http.StatusInternalServerError, "XML not valid: %v", validationError.Errors)
		} else if errors.Is(err, validator.ErrBadSyntax) {
			c.String(http.StatusInternalServerError, "XML not valid: %v", err)
		} else {
			c.String(http.StatusInternalServerError, "XML validation failed: %s", err.Error())
		}
		return
	}

	filename := filepath.Join("upload", filepath.Base(fileHeader.Filename))
	if err := c.SaveUploadedFile(fileHeader, filename); err != nil {
		c.String(http.StatusInternalServerError, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded and validated successfully", fileHeader.Filename)
}
