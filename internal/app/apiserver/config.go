package apiserver

type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config {
		BindAddress: "localhost:8080",
		LogLevel: "debug",
	}
}
