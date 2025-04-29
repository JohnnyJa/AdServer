package server

import (
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/reader"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/redis"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/druid-repo"
	"github.com/JohnnyJa/AdServer/EventAggregator/internal/storage/writer"
)

type Config struct {
	AppConfig    *AppConfig     `toml:"app"`
	RedisConfig  *redis.Config  `toml:"redis"`
	ReaderConfig *reader.Config `toml:"reader"`
	DruidConfig  *druid.Config  `toml:"druid"`
	WriterConfig *writer.Config `toml:"writer"`
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
		RedisConfig:  redis.NewConfig(),
		ReaderConfig: reader.NewConfig(),
	}
}
