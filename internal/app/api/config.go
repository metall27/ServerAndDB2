package api

import "github.com/metall27/ServerAndDB2/storage"

//General instance for API server of REST app

type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger level
	LoggerLevel string `toml:"logger_level"`
	//Storage configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
