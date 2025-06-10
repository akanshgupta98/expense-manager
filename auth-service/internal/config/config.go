package config

import (
	"os"

	"github.com/akanshgupta98/go-logger"
)

type Config struct {
	ServerConfig ServerConfig
	LoggerConfig logger.LogCfg
	SecretKey    string
	AMQPConfig   AMQPConfig
}

type AMQPConfig struct {
	Url            string
	PublishExhange string
}

type ServerConfig struct {
	WebPort string
}

func New() Config {

	serverCfg := ServerConfig{
		WebPort: "80",
	}

	loggerCfg := logger.LogCfg{
		Env:    os.Getenv("ENV"),
		Writer: os.Stdout,
	}

	amqp := AMQPConfig{
		Url:            os.Getenv("AMQP-URL"),
		PublishExhange: os.Getenv("AUTH-EXCHANGE"),
	}

	cfg := Config{
		ServerConfig: serverCfg,
		LoggerConfig: loggerCfg,
		SecretKey:    os.Getenv("JWT_Secret"),
		AMQPConfig:   amqp,
	}
	return cfg

}
