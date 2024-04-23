package main

import (
	"fmt"

	"github.com/HouseCham/cerebro/api/routes"
	"github.com/HouseCham/cerebro/internal/config"
	"github.com/HouseCham/cerebro/internal/log"
	"github.com/HouseCham/cerebro/api/core/grpc"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	// Setting up config file
	log.Logger.Infoln("Setting up config file")
	err := config.GetConfigFile()
	if err != nil {
		panic(err)
	}

	// Create a new Fiber instance
	log.Logger.Infoln("Setting up Fiber instance with CORS middleware and routes")
	app := fiber.New()
	// Setting up CORS middleware
	app.Use((cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE, PATCH",
	})))

	// setting up all routes
	routes.SetAllRoutes(app)

	// Setting up grpc connections
	log.Logger.Infoln("Setting up gRPC connections")
	grpc.SetUpGrpcConnections()


	// Start Fiber server
	log.Logger.Printf("Server is running on http://%s:%d", config.ConfigFile.App.Host, config.ConfigFile.App.Port)
	app.Listen(fmt.Sprintf(":%d", config.ConfigFile.App.Port))
}