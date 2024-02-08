package telegram

import (
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot       *tgbotapi.BotAPI
	pocket    *pocket.Client
	redirect  string
	tokenRepo repository.TokenRepository
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirect string, tokenRepo repository.TokenRepository) *Bot {
	return &Bot{bot: bot, pocket: pocketClient, redirect: redirect, tokenRepo: tokenRepo}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := b.initUpdatesChan()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) initUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u), nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand() {
				b.handleCommand(update.Message)
				continue
			}
			b.handleMsg(update.Message)

		}
	}
}
