package telegram

import (
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	pocket      *pocket.Client
	redirectURL string
	tokenRepo   repository.TokenRepo
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectURL string, tokenRepo repository.TokenRepo) *Bot {
	return &Bot{
		bot:         bot,
		pocket:      pocketClient,
		redirectURL: redirectURL,
		tokenRepo:   tokenRepo,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMsg(update.Message)
	}

	return nil
}
