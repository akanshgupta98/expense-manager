package config

import "github.com/akanshgupta98/go-logger/v2"

type Config struct {
	ServerCfg ServerCfg
	LoggerCfg logger.LogCfg
	JWTSecret string
}

type ServerCfg struct {
	WebPort string
}
