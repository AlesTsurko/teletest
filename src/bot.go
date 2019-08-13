package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func initBot(config *Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return nil, err
	}

	hook := tgbotapi.NewWebhook(config.Endpoint + bot.Token)

	_, err = bot.SetWebhook(hook)
	if err != nil {
		return bot, err
	}

	return bot, nil
}
