package server

import "github.com/JohnnyJa/AdServer/PackageService/internal/repository"

type Config struct {
	AppConfig      *AppConfig         `toml:"app"`
	PostgresConfig *repository.Config `toml:"postgres"`
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
