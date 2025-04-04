package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) HandleUpdate(update *tgbotapi.Update) {
	if update.CallbackQuery != nil {
		invoiceCommander.handleCallbackQuery(update.CallbackQuery)
		return
	}
	if update.Message != nil {
		invoiceCommander.handleMessage(update.Message)
		return
	}
}
