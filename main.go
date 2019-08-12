package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jinzhu/configor"
)

const (
	configPath = "config.yml"
)

// Config represents configuration of the service.
type Config struct {
	APIToken string
	Endpoint string
	Port     int `default:"8080"`
}

func initBot(config *Config) error {
	bot, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(config.Endpoint + bot.Token))
	if err != nil {
		return err
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.Port), nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

	return nil
}

func startServer(config *Config) error {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":" + strconv.Itoa(config.Port))

	return nil
}

func main() {
	config := Config{}
	configor.Load(&config, configPath)

	// err := initBot(&config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := startServer(&config)
	if err != nil {
		log.Fatal(err)
	}

}
