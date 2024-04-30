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
	log.Logger.Infoln("Finished CallInsertCostumer in grpc client")
	return model.HttpResponse{
		Data:         grpcResponse.NewId,
		HasError:     grpcResponse.HasError,
		ErrorMessage: grpcResponse.ErrorMessage,
		StatusCode:   utils.GrpcStatusToFiber(grpcResponse.StatusCode),
	}
}

// CallGetCustomerById is a function that calls the grpc CLIENT to get a customer by id
func CallGetCustomerById(id uint32) model.HttpResponse {
	log.Logger.Infoln("Initializing CallGetCustomerById in grpc client")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call grpc server
	grpcResponse, err := customerGrpcClient.GetCustomer(ctx, &GetCustomerQuery{Id: id})
	if err != nil {
		log.Logger.Errorf("There was an error trying to get customer by id: %v", err)
		return model.HttpResponse{
			Data:         nil,
			HasError:     true,
			ErrorMessage: fmt.Sprintf("There was an error with grpc server: %v", err),
			StatusCode:   fiber.StatusInternalServerError,
		}
	}
	log.Logger.Infoln("Finished CallGetCustomerById in grpc client")
	return model.HttpResponse{
		Data: &model.Customer{
			ID:          grpcResponse.Id,
			FirstName:   grpcResponse.FirstName,
			SecondName:  grpcResponse.SecondName,
			LastNameP:   grpcResponse.LastNameP,
			LastNameM:   grpcResponse.LastNameM,
			PhoneNumber: grpcResponse.PhoneNumber,
			Email:       grpcResponse.Email,
		},
		HasError:     grpcResponse.HasError,
		ErrorMessage: grpcResponse.ErrorMessage,
		StatusCode:   utils.GrpcStatusToFiber(grpcResponse.StatusCode),
	}
}
