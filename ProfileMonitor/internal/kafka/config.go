package kafka

type Config struct {
	Brokers []string `toml:"brokers"`
	Topic   string   `toml:"topic"`
}
