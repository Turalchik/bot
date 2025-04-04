package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) generateKeyboardMessage(callback *tgbotapi.CallbackQuery, isNextInvoices bool) {
	if isNextInvoices {
		invoiceCommander.cursor += invoiceCommander.limit
	} else {
		invoiceCommander.cursor -= invoiceCommander.limit
	}
	arrOfInvoices, err := invoiceCommander.invoiceService.List(invoiceCommander.cursor, invoiceCommander.limit)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка! "+err.Error()))
		return
	}

	outTxt := "Предыдущая порция счетов:\n\n"
	if isNextInvoices {
		outTxt = "Следующая порция счетов:\n\n"
	}
	outTxt += arrOfInvoices2Txt(arrOfInvoices)

	editMsg := tgbotapi.NewEditMessageText(callback.Message.Chat.ID, callback.Message.MessageID, outTxt)
	editMsg.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			firstRowForKeyboardMarkup(invoiceCommander.cursor, invoiceCommander.limit, invoiceCommander.invoiceService.GetSize()),
			{tgbotapi.NewInlineKeyboardButtonData("остановить показ", "stop")},
		},
	}

	invoiceCommander.bot.Send(editMsg)
}
