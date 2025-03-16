package main

import (
	"github.com/Turalchik/bot/internal/app/commands"
	"github.com/Turalchik/bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	commander := commands.NewCommander(bot, product.NewService())

	for update := range updates {
		commander.HandleUpdate(&update)
	}
}
