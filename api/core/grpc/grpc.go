package grpc

import customerService "github.com/HouseCham/cerebro/api/core/grpc/customerService"

// SetUpGrpcConnections is a function that sets up all the gRPC connections to all microservices
func SetUpGrpcConnections() {
	customerService.SetUpCostumerClient()
}