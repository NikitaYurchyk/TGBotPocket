package main

import (
	"github.com/NikitaYurchyk/TGPocket/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6984649114:AAGcKYIfSrVh23QZeUQxfCz7iVuW2ZjPWL8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient("110418-0db4c582cf23e23fcfa354a")
	if err != nil {
		log.Panic(err)
	}

	tgBot := telegram.NewBot(bot, pocketClient, "http://localhost")
	err = tgBot.Start()
	if err != nil {
		log.Fatal(err)
	}

}
