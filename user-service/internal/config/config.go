package config

import (
	"os"

	logger "github.com/akanshgupta98/go-logger/v2"
)

func New() Config {
	return Config{
		ServerCfg: ServerCfg{
			WebPort: "80",
		},
		LoggerCfg: logger.LogCfg{
			Env:    os.Getenv("ENV"),
			Writer: os.Stdout,
		},
		DBConfig: DBConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}
