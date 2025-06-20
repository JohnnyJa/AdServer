package app

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Name                      string        `toml:"name"`
	LogLevel                  string        `toml:"log_level"`
	ProfilesClientConfig      *ClientConfig `toml:"profiles"`
	ProfilesStateClientConfig *ClientConfig `toml:"profile_state"`
	*ServerConfig             `toml:"server"`
	*SemanticTargetingConfig  `toml:"semantic_targeting"`
}

type ServerConfig struct {
	Port string `toml:"port"`
}

type ClientConfig struct {
	Address string `toml:"address"`
}

type SemanticTargetingConfig struct {
	Path          string  `toml:"path"`
	MinSimilarity float64 `toml:"min_similarity"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadConfig() error {
	env := os.Getenv("APP_ENV")
	conf := os.Getenv("APP_CONF_PATH")
	if env == "" {
		env = "local"
	}

	_, err := toml.DecodeFile(conf+"config."+env+".toml", c)
	if err != nil {
		return err
	}
	return nil
}
