package grpc

import (
	"fmt"

	"github.com/HouseCham/cerebro/internal/config"
	"github.com/HouseCham/cerebro/internal/log"
	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var customerGrpcClient CustomerClient

// SetUpCostumerClient is a function that sets up the gRPC connection to the costumer service
func SetUpCostumerClient() {
	// Setting up costumer gRPC connection
	grpcConn, err := googleGrpc.Dial(fmt.Sprintf("%s:%d", config.ConfigFile.Microservices.CustomerService.Host, config.ConfigFile.Microservices.CustomerService.Port), googleGrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Logger.Fatalf("Failed to dial gRPC connection: %v", err)
	}

	// Set up costumer client
	customerGrpcClient = NewCustomerClient(grpcConn)
}
