package rest

import (
	"fmt"
	"iis_server/validator"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleXMLUpload(c *gin.Context) {
	// TODO maybe extract to global
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

	switch parts[len(parts)-1] {
	case "xsd":
		if err := validator.ValidateWithXSD(data, xsdFilePath); err != nil {
			// TODO match errors to return ok, bad xml
			c.String(http.StatusInternalServerError, "XML validation failed: %s", err.Error())
			return
		}
	case "rng":
		if err := validator.ValidateWithRNG(data, rngFilePath); err != nil {
			// TODO match errors to return ok, bad xml
			c.String(http.StatusInternalServerError, "XML validation failed: %s", err.Error())
			return
		}
	default:
		c.String(http.StatusInternalServerError, "Invalid validation: %s, %s", parts[len(parts)-1], err.Error())
	}

	filename := filepath.Join("upload", filepath.Base(fileHeader.Filename))
	if err := c.SaveUploadedFile(fileHeader, filename); err != nil {
		c.String(http.StatusInternalServerError, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded and validated successfully", fileHeader.Filename)
}
