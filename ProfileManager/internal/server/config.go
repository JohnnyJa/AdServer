package server

import (
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/kafka"
)

type Config struct {
	AppConfig   *AppConfig    `toml:"app"`
	KafkaConfig *kafka.Config `toml:"kafka"`
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
	}
}
