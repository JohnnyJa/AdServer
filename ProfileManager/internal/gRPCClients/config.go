package gRPCClients

import "time"

type Config struct {
	Address string        `toml:"address"`
	Timeout time.Duration `toml:"timeout"`
}

func NewConfig() Config {
	return Config{}
}
