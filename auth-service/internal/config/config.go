package config

import (
	"os"

	"github.com/akanshgupta98/go-logger"
)

type Config struct {
	ServerConfig ServerConfig
	LoggerConfig logger.LogCfg
}

type ServerConfig struct {
	WebPort string
}

func New() *Config {

	// serverCfg := ServerConfig{
	// 	WebPort: "8082",
	// }
	// loggerCfg :=
	return &Config{
		ServerConfig: ServerConfig{
			WebPort: "8082",
		},
		LoggerConfig: logger.LogCfg{
			Env:    logger.DEV_ENV,
			Writer: os.Stdout,
		},
	}
}
