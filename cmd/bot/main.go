package main

import (
	"github.com/NikitaYurchyk/TGPocket/pkg/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6984649114:AAGcKYIfSrVh23QZeUQxfCz7iVuW2ZjPWL8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	tgBot := telegram.NewBot(bot)
	err = tgBot.Start()
	if err != nil {
		log.Fatal(err)
	}

}
