package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/HouseCham/cerebro/api/core/controllers"
)

// SetAllRoutes sets all routes for the API Endpoints
func SetAllRoutes(app *fiber.App) {
	// Customer routes
	customerRoutes := app.Group("/api/v1/customer")
	customerRoutes.Post("/", controllers.InsertNewCustomer)
}