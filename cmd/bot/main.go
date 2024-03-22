package main

import (
	"errors"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository"
	"github.com/NikitaYurchyk/TGPocket/pkg/repository/db_bolt"
	"github.com/NikitaYurchyk/TGPocket/pkg/server"
	"github.com/NikitaYurchyk/TGPocket/pkg/telegram"
	"github.com/boltdb/bolt"
	"github.com/joho/godotenv"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	pocketClient, err := pocket.NewClient(os.Getenv("CONSUMER_KEY"))
	if err != nil {
		log.Panic(err)
	}

	db, err := bolt.Open("bot.db", 0777, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Batch(func(tx *bolt.Tx) error {
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

	tr := db_bolt.NewTokenStorage(db)

	tgBot := telegram.NewBot(bot, pocketClient, "http://localhost/", tr)
	redirServer := server.NewAuthServer("https://t.me/TGPocketProjectBot", tr, pocketClient)

	go func() {
		err = redirServer.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err := tgBot.Start(); err != nil {
		log.Fatal(err)
	}
}
