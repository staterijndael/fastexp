package store

// Config ...
type Config struct {
	DatabaseURL string `toml:"databaseUrl"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
