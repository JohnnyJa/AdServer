package server

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/redis"
)

type Config struct {
	AppConfig   *AppConfig    `toml:"app"`
	RedisConfig *redis.Config `toml:"redis"`
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
		RedisConfig: redis.NewConfig(),
	}
}
