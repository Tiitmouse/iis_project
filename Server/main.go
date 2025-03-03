package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", handleXMLUpload)
	r.Run(":8088")
}

func handleXMLUpload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get file err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "File %s uploaded successfully", file.Filename)
}
