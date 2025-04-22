package store

type Config struct {
	ConnectionString string `toml:"connection_string"`
}

func NewConfig() *Config {
	return &Config{
		ConnectionString: "",
	}
}
