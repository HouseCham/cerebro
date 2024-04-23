package model

// Config struct is used to read the config.json file
type Config struct {
	App           ServiceConnection `mapstructure:"app"`
	Microservices Microservices     `mapstructure:"microservices"`
	JWT           JWT               `mapstructure:"jwt"`
}

// ServiceConnection struct is used to read the application configuration
type ServiceConnection struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

// Microservices struct is used to read the microservices configuration
type Microservices struct {
	CustomerService ServiceConnection `mapstructure:"customerService"`
}

// JWT struct is used to read the JWT configuration
type JWT struct {
	Secret string `mapstructure:"secret"`
}
