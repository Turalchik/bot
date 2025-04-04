package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) handleCallbackQuery(callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "next_invoices":
		invoiceCommander.generateKeyboardMessage(callback, true)
	case "previous_invoices":
		invoiceCommander.generateKeyboardMessage(callback, false)
	default:
		invoiceCommander.bot.Send(tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Показ закончен")
		invoiceCommander.bot.Send(msg)
	}
}
