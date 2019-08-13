package main

import (
	"log"

	"github.com/jinzhu/configor"
)

const (
	configPath = "config.yml"
)

func main() {
	config := Config{}
	configor.Load(&config, configPath)

	bot, err := initBot(&config)
	checkError(&err)

	err = startServer(&config, bot)
	checkError(&err)
}

// Config represents configuration of the service.
type Config struct {
	APIToken string
	Endpoint string
	Port     int `default:"8000"`
}

func checkError(err *error) {
	if *err != nil {
		log.Fatal(err)
	}
}
