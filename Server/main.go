package main

import (
	"fmt"
	"iis_server/handlers"
	"iis_server/storage"
	"iis_server/utils"
	"iis_server/xmlrpcserver"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gorilla_rpc "github.com/gorilla/rpc"
)

func main() {
	utils.LoadEnv()

	// Initialize token store with custom file path (optional)
	storage.SetTokenStoreFilePath("tokens.json")

	contactStore := storage.NewContactStore()
	handlers.SetContactStore(contactStore)

	storage.SeedStore(contactStore)

	fmt.Println("Setting up Gorilla XML-RPC server...")
	xmlrpc := gorilla_rpc.NewServer()
	xmlCodec := xml.NewCodec()
	xmlrpc.RegisterCodec(xmlCodec, "text/xml")
	err := xmlrpc.RegisterService(new(xmlrpcserver.WeatherService), "")
	if err != nil {
		log.Fatalf("Failed to register XML-RPC service: %v", err)
	}
	xmlrpcAddr := ":8089"
	http.Handle("/RPC2", xmlrpc)

	go func() {
		fmt.Printf("Starting Gorilla XML-RPC listener on %s/RPC2\n", xmlrpcAddr)
		err := http.ListenAndServe(xmlrpcAddr, nil)
		if err != nil {
			log.Printf("FATAL: XML-RPC server ListenAndServe error on %s: %v\n", xmlrpcAddr, err)
		}
	}()

	fmt.Println("Setting up Gin server...")
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
		api.POST("/logout", handlers.JWTMiddleware(), handlers.LogoutHandler)
		api.POST("/refresh", handlers.JWTMiddleware(), handlers.RefreshTokenHandler)
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

	ginPort := ":8088"
	fmt.Printf("Starting Gin server on port %s\n", ginPort)
	ginErr := r.Run(ginPort)
	if ginErr != nil {
		log.Fatalf("Error starting Gin server: %v", ginErr)
	}
}
