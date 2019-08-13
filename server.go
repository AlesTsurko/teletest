package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func startServer(config *Config) error {
	router := gin.Default()

	router.POST("/"+config.APIToken, replyRout)

	router.Run(":" + strconv.Itoa(config.Port))

	return nil
}

func replyRout(c *gin.Context) {
	teleRequest, err := deserializeRequest(c)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(teleRequest)
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

// TelegramRequest is serialized JSON bot request.
type TelegramRequest struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is Telegram's message type.
type Message struct {
	Text string `json:"text"`
}
