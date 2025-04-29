package writer

import "time"

type Config struct {
	NumOfWriters int           `toml:"num_of_writers"`
	FlushTimeout time.Duration `toml:"flush_timeout"`
	MaxSize      int           `toml:"max_size"`
}

func NewConfig() *Config {
	return &Config{
		NumOfWriters: 10,
	}
}
