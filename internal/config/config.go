package config

import (
	"os"
	"strconv"

	"github.com/HouseCham/cerebro/internal/model"
)

var ConfigFile model.Config

// getConfigFile reads the config.json file and returns a Config struct
func GetConfigFile() error {
	// Parse the cerebro port from the environment variable
	cerebroPort, err := strconv.ParseUint(os.Getenv("CEREBRO_PORT"), 10, 16)
	if err != nil {
		return err
	}
	// Parse the customer service port from the environment variable
	customerServicePort, err := strconv.ParseUint(os.Getenv("CUSTOMER_SERVICE_PORT"), 10, 16)
	if err != nil {
		return err
	}
	// Unmarshal the config file into a struct
	var configFile model.Config = model.Config{
		App: model.ServiceConnection{
			Host: os.Getenv("CEREBRO_HOST"),
			Port: uint16(cerebroPort),
		},
		Microservices: model.Microservices{
			CustomerService: model.ServiceConnection{
				Host: os.Getenv("CUSTOMER_SERVICE_HOST"),
				Port: uint16(customerServicePort),
			},
		},
		JWT: model.JWT{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}


	ConfigFile = configFile
	return nil
}
