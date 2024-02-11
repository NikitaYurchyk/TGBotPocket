package telegram

import (
	"context"
	"fmt"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) initAuthMsg(message *tgbotapi.Message) error {
	authLink, _ := b.generateAuthLink(message.Chat.ID)
	msgText := fmt.Sprintf(REPLY, authLink)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) generateAuthLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectURL(chatID)
	token, err := b.pocket.GetRequestToken(context.Background(), b.redirectURL)
	if err != nil {
		return "", err
	}

	if err := b.tokenRepo.Save(chatID, token, repository.RequestTokens); err != nil {
		return "", err
	}
	return b.pocket.GetAuthorizationURL(token, redirectUrl)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepo.Get(chatID, repository.AccessTokens)
}
