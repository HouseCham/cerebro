package grpc

import (
	context "context"
	"fmt"
	"time"

	"github.com/HouseCham/cerebro/internal/log"
	"github.com/HouseCham/cerebro/internal/model"
	"github.com/HouseCham/cerebro/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

// CallInsertCostumer is a function that calls the grpc CLIENT to insert a new costumer
func CallInsertCostumer(command *InsertCustomerCommand) model.HttpResponse {
	log.Logger.Infoln("Initializing CallInsertCostumer in grpc client")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call grpc server
	grpcResponse, err := customerGrpcClient.InsertCustomer(ctx, command)
	if err != nil {
		log.Logger.Errorf("There was an error trying to insert customer: %v", err)
		return model.HttpResponse{
			Data:         0,
			HasError:     true,
			ErrorMessage: fmt.Sprintf("There was an error with grpc server: %v", err),
			StatusCode:   fiber.StatusInternalServerError,
		}
	}
	return model.HttpResponse{
		Data:         grpcResponse.NewId,
		HasError:     grpcResponse.HasError,
		ErrorMessage: grpcResponse.ErrorMessage,
		StatusCode:   utils.GrpcStatusToFiber(grpcResponse.StatusCode),
	}
}
