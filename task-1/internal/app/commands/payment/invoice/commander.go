package invoice

import (
	service "github.com/Turalchik/bot/task-1/internal/service/payment/invoice"
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
	bot                   *tgbotapi.BotAPI
	invoiceService        service.InvoiceService
	cursor                uint64
	limit                 uint64
	lastKeyboardMessageID int
}

func NewInvoiceCommander(bot *tgbotapi.BotAPI, service service.InvoiceService) InvoiceCommander {
	return &DummyInvoiceCommander{
		bot:                   bot,
		invoiceService:        service,
		lastKeyboardMessageID: -1,
	}
}
