package invoice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (invoiceCommander *DummyInvoiceCommander) HandleUpdate(update *tgbotapi.Update) {
	if update.CallbackQuery != nil {
		invoiceCommander.handleCallbackQuery(update.CallbackQuery)
		return
	}

	if update.Message != nil {
		if invoiceCommander.lastKeyboardMessageID >= 0 {
			invoiceCommander.deleteKeyboard(update.Message.Chat.ID, invoiceCommander.lastKeyboardMessageID)
		}
		invoiceCommander.handleMessage(update.Message)
		return
	}
}
