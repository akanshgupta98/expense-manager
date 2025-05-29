package config

import (
	"os"

	logger "github.com/akanshgupta98/go-logger/v2"
)

func New() Config {
	return Config{
		ServerCfg: ServerCfg{
			WebPort: "8083",
		},
		LoggerCfg: logger.LogCfg{
			Env:    os.Getenv("ENV"),
			Writer: os.Stdout,
		},
	}
}
