package invoice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"unicode"
)

func (invoiceCommander *DummyInvoiceCommander) Edit(inputMsg *tgbotapi.Message) {
	stringOfArguments := strings.TrimFunc(inputMsg.CommandArguments(), unicode.IsSpace)
	args := strings.Split(stringOfArguments, " ")

	if len(args) != 4 {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID,
			"Ошибка! Формат: /edit__payment__invoice <ID> <новый счёт> <новая сумма> <новая валюта>"))
		return
	}

	ID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! ID должен быть натуральным числом"))
		return
	}

	invoice, err := getInvoice(args[1:])
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error()))
		return
	}

	err = invoiceCommander.invoiceService.Update(ID, *invoice)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error()))
		return
	}

	msg := fmt.Sprintf("Счет успешно обновлён! №%s на сумму %.2f %s\n", invoice.Number, invoice.Amount, invoice.Currency)
	invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, msg))
}
