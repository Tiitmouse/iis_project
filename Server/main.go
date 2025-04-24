package main

import (
	"iis_server/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.POST("/login", handlers.LoginHandler)
		api.POST("/refresh", handlers.RefreshTokenHandler)
	}

	r.POST("/upload/xsd", handlers.HandleXMLUpload)
	r.POST("/upload/rng", handlers.HandleXMLUpload)

	r.GET("/weather", handlers.GetWeatherByCity)

	r.Run(":8088")
}
