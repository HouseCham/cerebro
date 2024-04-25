package controllers

import (
	"fmt"

	customerService "github.com/HouseCham/cerebro/api/core/grpc/customerService"
	"github.com/HouseCham/cerebro/internal/log"
	"github.com/HouseCham/cerebro/internal/model"
	"github.com/gofiber/fiber/v3"
)

// InsertNewCostumer is a controller that inserts a new costumer
func InsertNewCustomer(c fiber.Ctx) error {
	log.Logger.Infoln("InsertNewCustomer controller invoked")

	var request customerService.InsertCustomerCommand
	// parsing request body from JSON to struct
	if err := c.Bind().JSON(&request); err != nil {
		log.Logger.Warnf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(model.HttpResponse{
			HasError:   true,
			ErrorMessage:      fmt.Sprintf("Error parsing request body: %v", err),
			StatusCode: fiber.StatusBadRequest,
			Data:       nil,
		})
	}

	// calling gRPC service
	serverResponse := customerService.CallInsertCostumer(&request)

	return c.Status(fiber.StatusOK).JSON(serverResponse)
}