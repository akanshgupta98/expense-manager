package config

type Config struct {
	ServerConfig ServerConfig
}

type ServerConfig struct {
	WebPort string
}

func New() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			WebPort: "8082",
		},
	}
}
