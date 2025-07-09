package invoice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"unicode"
)

func (invoiceCommander *DummyInvoiceCommander) List(inputMsg *tgbotapi.Message) {
	arg := strings.TrimFunc(inputMsg.CommandArguments(), unicode.IsSpace)
	if arg == "" {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! Формат: /list__payment__invoice <limit>"))
		return
	}

	limit, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! limit должен быть натуральным числом"))
		return
	}

	invoiceCommander.cursor = 0
	invoiceCommander.limit = limit

	arrOfInvoices, err := invoiceCommander.invoiceService.List(invoiceCommander.cursor, invoiceCommander.limit)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! "+err.Error()))
		return
	}

	outTxt := "Следующая порция счетов:\n\n"
	outTxt += arrOfInvoices2Txt(arrOfInvoices)
	outMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, outTxt)

	outMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		firstRowForKeyboardMarkup(invoiceCommander.cursor, invoiceCommander.limit, invoiceCommander.invoiceService.GetSize()),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("остановить показ", "stop"),
		),
	)

	sentMsg, _ := invoiceCommander.bot.Send(outMsg)
	invoiceCommander.lastKeyboardMessageID = sentMsg.MessageID
}
