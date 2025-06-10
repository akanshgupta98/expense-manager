package config

import "github.com/akanshgupta98/go-logger/v2"

type Config struct {
	ServerCfg  ServerCfg
	LoggerCfg  logger.LogCfg
	DBConfig   DBConfig
	AMQPConfig AMQPConfig
}

type AMQPConfig struct {
	URL              string
	ConsumeExchanges []string
}
type ServerCfg struct {
	WebPort string
}

type DBConfig struct {
	DSN string
}
