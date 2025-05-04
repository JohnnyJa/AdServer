package worker

import "time"

type Config struct {
	Delay time.Duration `toml:"delay"`
}
