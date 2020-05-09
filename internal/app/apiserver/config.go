package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bindAddr"`
	LogLevel    string `toml:"logLevel"`
	DatabaseURL string `toml:"databaseUrl`
	SessionKey  string `toml:"sessionKey"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
