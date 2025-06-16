package config

import (
	"os"

	"github.com/akanshgupta98/go-logger/v2"
)

func New() Config {

	cfg := Config{
		ServerCfg: ServerCfg{
			WebPort: "80",
		},
		LoggerCfg: logger.LogCfg{
			Writer: os.Stdout,
			Env:    os.Getenv("ENV"),
		},
	}
	return cfg

}
