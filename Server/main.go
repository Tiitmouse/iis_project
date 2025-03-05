package main

import (
	"iis_server/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", rest.HandleXMLUpload)
	r.Run(":8088")
}
