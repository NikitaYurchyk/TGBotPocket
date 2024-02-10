package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/url"
)

const (
	START        = "start"
	REPLY        = "Hey! Before adding bookmarks, i need to receive an access to your pocket.\n%s"
	UNKNOWN      = "Unknown command"
	AUTHORIZED   = "Already authorized. Now, you can send your links, I will save them!"
	UNAUTHORIZED = "Not authorized. Try to use command /start"
	INVALID_LINK = "Invalid link"
	ERROR_SAVE   = "Link can not be save. Please, try again."
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
		fmt.Println("No access token found")
		return b.initAuthMsg(message)
	}
	token, _ := b.getAccessToken(message.Chat.ID)
	fmt.Println(token)

	msg := tgbotapi.NewMessage(message.Chat.ID, AUTHORIZED)

	_, err = b.bot.Send(msg)

	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleMsg(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = INVALID_LINK
		_, err = b.bot.Send(msg)
		return err
	}
	accessToken, err := b.getAccessToken(message.Chat.ID)

	fmt.Println("ACCESS TOKEN", accessToken, err)
	if err != nil {
		return b.initAuthMsg(message)
	}
	fmt.Println(accessToken)
	if err := b.pocket.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = ERROR_SAVE
		fmt.Println(err)
		_, err = b.bot.Send(msg)
		return err
	}

	_, err = b.bot.Send(msg)

	if err != nil {
		_ = fmt.Errorf("%s\n", err)
	}
	return nil
}
