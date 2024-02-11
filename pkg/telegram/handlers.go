package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

const (
	START      = "start"
	REPLY      = "Hey! Before adding bookmarks, i need to receive an access to your pocket.\n%s"
	UNKNOWN    = "Unknown command"
	AUTHORIZED = "Already authorized. Now, you can send your links, I will save them!"
	SAVED      = "Link was saved"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case START:
		return b.handleStart(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, UNKNOWN)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthMsg(message)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, AUTHORIZED)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleMsg(message *tgbotapi.Message) error {
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthMsg(message)
	}

	if err := b.handleSaving(message, accessToken); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, SAVED)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleSaving(message *tgbotapi.Message, accessToken string) error {
	if err := b.validateURL(message.Text); err != nil {
		return err
	}

	if err := b.pocket.Add(context.Background(), pocket.AddInput{
		URL:         message.Text,
		AccessToken: accessToken,
	}); err != nil {
		fmt.Println("Not Added")
		return err
	}

	return nil
}

func (b *Bot) validateURL(text string) error {
	_, err := url.ParseRequestURI(text)
	return err
}
