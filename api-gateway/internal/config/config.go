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
			Env:    os.Getenv("ENV"),
			Writer: os.Stdout,
		},
	}
	return cfg

}
