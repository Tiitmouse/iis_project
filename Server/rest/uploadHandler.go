package rest

import (
	"fmt"
	"iis_server/validator"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleXMLUpload(c *gin.Context) {
	// TODO maybe extract to global
	xsdFilePath := "./schemas/exampleXSD.xsd"

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

	if err := validator.ValidateWithXSD(data, xsdFilePath); err != nil {
		// TODO match errors to return ok, bad xml
		c.String(http.StatusInternalServerError, "XML validation failed: %s", err.Error())
		return
	}

	filename := filepath.Join("upload", filepath.Base(fileHeader.Filename))
	if err := c.SaveUploadedFile(fileHeader, filename); err != nil {
		c.String(http.StatusInternalServerError, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded and validated successfully", fileHeader.Filename)
}
