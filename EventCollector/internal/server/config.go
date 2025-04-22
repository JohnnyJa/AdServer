package server

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/store"
)

type Config struct {
	AppConfig   *AppConfig    `toml:"app"`
	RedisConfig *store.Config `toml:"store"`
}

type AppConfig struct {
	Name     string `toml:"name"`
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		AppConfig: &AppConfig{
			Name:     "",
			Port:     "",
			LogLevel: "info",
		},
		RedisConfig: store.NewConfig(),
	}
}
