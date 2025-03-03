package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", handleXMLUpload)
	r.Run(":8088")
}

func handleXMLUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "bokic"})
}
