package apiserver

import "github.com/Bratello-0/identification-api/internal/app/store"

type Config struct {
	Address  string `toml:"address"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Address:  "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
