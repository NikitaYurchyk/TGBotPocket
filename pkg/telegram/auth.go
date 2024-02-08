package telegram

import (
	"context"
	"fmt"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
)

func (b *Bot) generateAuthLink(chatId int64) (string, error) {
	redirectID := b.generateRedirectURL(chatId, b.redirect)
	requestToken, err := b.pocket.GetRequestToken(context.Background(), b.redirect)
	if err != nil {
		return "Error", err
	}

	if err := b.tokenRepo.Save(chatId, requestToken, repository.RequestTokens); err != nil {
		return "", err
	}

	return b.pocket.GetAuthorizationURL(requestToken, redirectID)
}

func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
	fmt.Println("START")
	fmt.Printf("%s?chat_id=%d\n", redirectURL, chatID)
	fmt.Println("OVER")
	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
}
