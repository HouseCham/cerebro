package utils

import (
	"google.golang.org/grpc/codes"
	"github.com/gofiber/fiber/v3"
)

// grpcStatusToFiber is a function that converts gRPC status codes to Fiber status codes
func GrpcStatusToFiber(statusCode uint32) uint32 {
	status := codes.Code(statusCode)
	switch status {
	case codes.OK:
		return fiber.StatusOK
	case codes.Canceled:
		return fiber.StatusRequestTimeout
	case codes.Unknown:
		return fiber.StatusInternalServerError
	case codes.InvalidArgument:
		return fiber.StatusBadRequest
	case codes.DeadlineExceeded:
		return fiber.StatusGatewayTimeout
	case codes.NotFound:
		return fiber.StatusNotFound
	case codes.AlreadyExists:
		return fiber.StatusConflict
	case codes.PermissionDenied:
		return fiber.StatusForbidden
	case codes.Unauthenticated:
		return fiber.StatusUnauthorized
	case codes.ResourceExhausted:
		return fiber.StatusTooManyRequests
	case codes.FailedPrecondition:
		return fiber.StatusPreconditionFailed
	case codes.Aborted:
		return fiber.StatusConflict
	case codes.OutOfRange:
		return fiber.StatusBadRequest
	case codes.Unimplemented:
		return fiber.StatusNotImplemented
	case codes.Internal:
		return fiber.StatusInternalServerError
	case codes.Unavailable:
		return fiber.StatusServiceUnavailable
	case codes.DataLoss:
		return fiber.StatusInternalServerError
	default:
		return fiber.StatusInternalServerError
	}
}