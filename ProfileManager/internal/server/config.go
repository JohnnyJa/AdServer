package server

import (
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/gRPCClients"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/profileStorage"
)

type Config struct {
	AppConfig           *AppConfig             `toml:"app"`
	ProfileClientConfig *gRPCClients.Config    `toml:"profile"`
	PackageClientConfig *gRPCClients.Config    `toml:"package"`
	StorageConfig       *profileStorage.Config `toml:"storage"`
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
