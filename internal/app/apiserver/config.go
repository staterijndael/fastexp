package apiserver

import "github.com/Oringik/fastexp/internal/app/store"

type Config struct {
	BindAddr string `toml:"bindAddr"`
	LogLevel string `toml:"logLevel"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
