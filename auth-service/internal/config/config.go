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

	serverCfg := ServerConfig{
		WebPort: "80",
	}

	loggerCfg := logger.LogCfg{
		Env:    os.Getenv("ENV"),
		Writer: os.Stdout,
	}

	cfg := Config{
		ServerConfig: serverCfg,
		LoggerConfig: loggerCfg,
	}
	return &cfg

}
