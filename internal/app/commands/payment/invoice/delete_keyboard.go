package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) deleteKeyboard(chatID int64, messageID int) {
	emptyMarkup := tgbotapi.NewEditMessageReplyMarkup(
		chatID, messageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}},
	)
	invoiceCommander.bot.Send(emptyMarkup)
	msg := tgbotapi.NewMessage(chatID, "Показ закончен")
	invoiceCommander.bot.Send(msg)
}
