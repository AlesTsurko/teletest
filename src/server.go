package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func startServer(config *Config, bot *tgbotapi.BotAPI) error {
	router := gin.Default()

	router.Use(setBot(bot))

	router.POST("/"+config.APIToken, replyRoute)

	router.Run(":" + strconv.Itoa(config.Port))

	return nil
}

func setBot(bot *tgbotapi.BotAPI) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("bot", bot)
		c.Next()
	}
}

func replyRoute(c *gin.Context) {
	teleRequest, err := deserializeRequest(c)
	if err != nil {
		log.Println(err)
		return
	}

	bot := c.MustGet("bot").(*tgbotapi.BotAPI)
	
	msg := initReplyMessage(teleRequest)

	bot.Send(msg)
}

func deserializeRequest(c *gin.Context) (*TelegramRequest, error) {
	decoder := json.NewDecoder(c.Request.Body)

	var deserialized TelegramRequest
	err := decoder.Decode(&deserialized)
	if err != nil {
		return nil, err
	}

	return &deserialized, nil
}

func initReplyMessage(teleRequest *TelegramRequest) tgbotapi.MessageConfig {
	text := "Robot say:\n" + teleRequest.Message.Text
	message := tgbotapi.NewMessage(teleRequest.Message.Chat.ID, text)
	message.ReplyToMessageID = teleRequest.Message.ID
	
	return message
}

// TelegramRequest is serialized JSON bot request.
type TelegramRequest struct {
	Message  Message `json:"message"`
}

// Message is Telegram's message type.
type Message struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Chat Chat `json:"chat"`
}

// Chat is Telegram's chat type.
type Chat struct {
	ID int64 `json:"id"`
}
