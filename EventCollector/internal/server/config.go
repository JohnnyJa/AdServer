package server

import (
	"github.com/JohnnyJa/AdServer/EventCollector/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/kafka"
)

type Config struct {
	AppConfig    *AppConfig                `toml:"app"`
	KafkaConfig  *kafka.Config             `toml:"kafka"`
	ClientConfig *grpcClients.ClientConfig `toml:"client"`
}

type AppConfig struct {
	Name     string `toml:"name"`
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
}

type ProducerConfig struct {
	Brokers []string `toml:"brokers"`
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
