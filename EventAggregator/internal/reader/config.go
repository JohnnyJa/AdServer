package reader

type Config struct {
	NumOfReaders int `toml:"num_of_readers"`
	BufferSize   int `toml:"buffer_size"`
}

func NewConfig() *Config {
	return &Config{
		NumOfReaders: 10,
		BufferSize:   1024,
	}
}
