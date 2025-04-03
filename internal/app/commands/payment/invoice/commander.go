package invoice

import (
	//model "github.com/Turalchik/bot/internal/model/payment"
	service "github.com/Turalchik/bot/internal/service/payment/invoice"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InvoiceCommander interface {
	HandleUpdate(update *tgbotapi.Update)
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type DummyInvoiceCommander struct {
	bot            *tgbotapi.BotAPI
	invoiceService service.InvoiceService
}

func NewInvoiceCommander(bot *tgbotapi.BotAPI, service service.InvoiceService) InvoiceCommander {
	return &DummyInvoiceCommander{
		bot:            bot,
		invoiceService: service,
	}
}
