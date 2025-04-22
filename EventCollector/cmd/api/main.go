package main

import (
	"github.com/BurntSushi/toml"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/server"
	"log"
	"os"
)

func main() {

	env := os.Getenv("APP_ENV")
	conf := os.Getenv("APP_CONF_PATH")
	if env == "" {
		env = "local"
	}

	config := server.NewConfig()

	_, err := toml.DecodeFile(conf+"config."+env+".toml", config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
