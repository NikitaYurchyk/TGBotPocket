package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	START   = "start"
	REPLY   = "Hey! Before adding bookmarks, i need to receive an access to your pocket.\n%s\n"
	UNKNOWN = "Unknown command"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, UNKNOWN)
	switch message.Command() {
	case START:
		b.handleStart(message)
		return nil
	default:
		_, err := b.bot.Send(msg)
		return err

	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	authLink, _ := b.generateAuthLink(message.Chat.ID)
	message.Text = START
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(REPLY, authLink))
	msg.ReplyToMessageID = message.MessageID
	newMsg, err := b.bot.Send(msg)

	if err != nil {
		_ = fmt.Errorf("%s\n", err)
	}
	log.Printf("%s\n", newMsg.Chat)
}

func (b *Bot) handleMsg(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	newMsg, err := b.bot.Send(msg)

	if err != nil {
		_ = fmt.Errorf("%s\n", err)
	}
	log.Printf("%s\n", newMsg.Chat)
}
