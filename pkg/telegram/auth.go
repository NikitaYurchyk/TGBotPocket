package telegram

import (
	"context"
	"errors"
	"fmt"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) initAuthMsg(message *tgbotapi.Message) error {
	authLink, _ := b.generateAuthLink(message.Chat.ID)
	fmt.Println("\n", authLink, "\n")
	msgText := fmt.Sprintf(REPLY, authLink)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) getAccessToken(chatID int64) (string, error) {
	return b.tokenRepo.Get(chatID, repository.AccessTokens)
}

func (b *Bot) generateAuthLink(chatId int64) (string, error) {
	redirectID := b.generateRedirectURL(chatId)
	fmt.Println("Redirect_ID:", redirectID)
	requestToken, err := b.pocket.GetRequestToken(context.Background(), b.redirect)
	if err != nil {
		return "Error", err
	}
	fmt.Println("RequestToken: ", requestToken)

	if err := b.tokenRepo.Save(chatId, requestToken, repository.RequestTokens); err != nil {
		return "", errors.New("Not saved")
	}
	fmt.Println("request token saved")

	check, err := b.getAccessToken(chatId)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Not found")
	}

	fmt.Println("CHECK", check)

	return b.pocket.GetAuthorizationURL(requestToken, redirectID)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirect, chatID)
}
