package profileStorage

import "time"

type Config struct {
	Interval time.Duration `toml:"interval"`
}

func NewConfig() *Config {
	return &Config{}
}
