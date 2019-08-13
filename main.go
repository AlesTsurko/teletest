package main

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
)

const defaultConfigPath = "config.yml"

func main() {
	config := Config{}
	configor.Load(&config, getConfigPath())

	bot, err := initBot(&config)
	checkError(&err)

	err = startServer(&config, bot)
	checkError(&err)
}

func getConfigPath() string {
	configPath, ok := os.LookupEnv("TELEBOT_CONFIG_PATH")
	if !ok {
		configPath = defaultConfigPath
	}

	return configPath
}

func checkError(err *error) {
	if *err != nil {
		log.Fatal(*err)
	}
}

// Config represents configuration of the service.
type Config struct {
	APIToken string
	Endpoint string
	Port     int `default:"8000"`
}
