package main

import (
	"github.com/Turalchik/bot/internal/app/commands/payment/invoice"
	service "github.com/Turalchik/bot/internal/service/payment/invoice"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	invoiceCommander := invoice.NewInvoiceCommander(bot, service.NewDummyInvoiceService())

	for update := range updates {
		invoiceCommander.HandleUpdate(&update)
	}
}
