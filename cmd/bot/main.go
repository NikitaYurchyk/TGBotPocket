package main

import (
	"errors"
	"fmt"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository/bolt"
	"github.com/NikitaYurchyk/TGPocket/pkg/server"
	"github.com/NikitaYurchyk/TGPocket/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"go.etcd.io/bbolt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6984649114:AAGcKYIfSrVh23QZeUQxfCz7iVuW2ZjPWL8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient("110418-0db4c582cf23e23fcfa354a")
	if err != nil {
		log.Panic(err)
	}

	db, err := bbolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Batch(func(tx *bbolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return errors.New("\nACCESS-BATCH NOT CREATED!\n")
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return errors.New("\nREQUEST-BATCH NOT CREATED!\n")
		}
		return nil
	})

	tr := bolt.NewTokenRepository(db)

	tgBot := telegram.NewBot(bot, pocketClient, "https://t.me/TGPocketProjectBot", tr)
	authServer := server.InitAuthServer(pocketClient, tr, "https://t.me/TGPocketProjectBot")

	go func() {
		err = tgBot.Start()
		if err != nil {
			fmt.Println("here")
			log.Fatal(err)
		}
	}()

	if err := authServer.Start(); err != nil {
		log.Fatal(err)
	}
}
