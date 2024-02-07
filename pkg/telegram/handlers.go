package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const START = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command")

	switch message.Command() {
	case START:
		msg.Text = "Bot started"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err

	}
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
