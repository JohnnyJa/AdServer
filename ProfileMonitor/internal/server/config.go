package server

import (
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/kafka"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/repository"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/internal/worker"
)

type Config struct {
	AppConfig      *AppConfig         `toml:"app"`
	PostgresConfig *repository.Config `toml:"postgres"`
	KafkaConfig    *kafka.Config      `toml:"kafka"`
	WorkerConfig   *worker.Config     `toml:"worker"`
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
		KafkaConfig: &kafka.Config{
			Brokers: []string{"localhost:9092"},
		},
	}
}
