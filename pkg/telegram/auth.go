package telegram

import (
	"context"
	"fmt"
)

func (b *Bot) generateAuthLink(chatId int64) (string, error) {
	redirectID := b.generateRedirectURL(chatId, b.redirect)
	requestToken, err := b.pocket.GetRequestToken(context.Background(), b.redirect)
	if err != nil {
		return "Error", err
	}
	return b.pocket.GetAuthorizationURL(requestToken, redirectID)
}

func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
}
