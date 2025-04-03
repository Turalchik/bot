package invoice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"unicode"
)

func (invoiceCommander *DummyInvoiceCommander) Get(inputMsg *tgbotapi.Message) {
	arg := strings.TrimFunc(inputMsg.CommandArguments(), unicode.IsSpace)
	if arg == "" {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! Формат: /get__payment__invoice <ID>"))
		return
	}

	ID, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! ID должен быть натуральным числом"))
		return
	}

	invoice, err := invoiceCommander.invoiceService.Describe(ID)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! "+err.Error()))
		return
	}

	msgText := fmt.Sprintf(
		"*Счет №%s*\n💰 *Сумма:* %.2f %s",
		invoice.Number, invoice.Amount, invoice.Currency,
	)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, msgText)
	msg.ParseMode = "Markdown"

	invoiceCommander.bot.Send(msg)
}
