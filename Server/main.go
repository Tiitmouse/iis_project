package main

import (
	"iis_server/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload/xsd", rest.HandleXMLUpload)
	r.POST("/upload/rng", rest.HandleXMLUpload)
	r.Run(":8088")
}
