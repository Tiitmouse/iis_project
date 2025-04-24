package main

import (
	"iis_server/rest"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	r.Use(cors.New(config))

	r.POST("/upload/xsd", rest.HandleXMLUpload)
	r.POST("/upload/rng", rest.HandleXMLUpload)
	r.Run(":8088")
}
