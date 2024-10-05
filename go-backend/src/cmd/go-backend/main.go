package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-backend/config"
	"go-backend/middleware"
	"go-backend/routes"
	"log"
)

func main () {
	log.Println("cmd/go-backend/main.go")
	// Load config
	config.LoadConfig()

	// Create a new fiber instance
	app := fiber.New() // app is a pointer
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// Middleware
	// log every request
	app.Use(middleware.Log)

	// Route
	routes.SetupRoutes(app)

	// Start the server on port 5000
	log.Printf("%#v\n", config.BackendConfig)
	err := app.Listen(":"+config.BackendConfig.ServerPort)
	if err != nil {
		panic(err)
	}
}