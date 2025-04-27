package main

import (
	"fmt"
	"iis_server/handlers"
	"iis_server/storage"
	"iis_server/utils"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()

	contactStore := storage.NewContactStore()
	handlers.SetContactStore(contactStore)

	storage.SeedStore(contactStore)

	r := gin.Default()

	config := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || strings.HasPrefix(origin, "wails://")

		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	r.GET("/weather", handlers.GetWeatherByCity)

	r.POST("/upload/xsd", handlers.HandleXMLUpload)
	r.POST("/upload/rng", handlers.HandleXMLUpload)

	api := r.Group("/api")
	{
		api.POST("/login", handlers.LoginHandler)
		api.POST("/refresh", handlers.RefreshTokenHandler)

		contacts := api.Group("/contacts")
		contacts.Use(handlers.JWTMiddleware())
		{
			contacts.POST("", handlers.CreateContact)
			contacts.GET("", handlers.GetAllContacts)
			contacts.GET("/:id", handlers.GetContactByID)
			contacts.PUT("/:id", handlers.UpdateContact)
			contacts.DELETE("/:id", handlers.DeleteContact)
		}
	}

	port := ":8088"
	fmt.Printf("Starting server on port %s\n", port)
	err := r.Run(port)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
