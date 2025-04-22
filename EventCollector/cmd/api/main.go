package main

import (
	"github.com/BurntSushi/toml"
	"github.com/JohnnyJa/AdServer/EventCollector/container/server"
	"log"
	"os"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	config := server.NewConfig()

	_, err := toml.DecodeFile("EventCollector/config/config."+env+".toml", config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}
}
