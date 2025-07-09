package invoice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"unicode"
)

func (invoiceCommander *DummyInvoiceCommander) New(inputMsg *tgbotapi.Message) {
	stringOfArguments := strings.TrimFunc(inputMsg.CommandArguments(), unicode.IsSpace)
	args := strings.Split(stringOfArguments, " ")

	invoice, err := getInvoice(args)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error()))
		return
	}
	ID, _ := invoiceCommander.invoiceService.Create(*invoice)

	msg := fmt.Sprintf("Счет создан! №%s на сумму %.2f %s\nID созданного счёта: %v", invoice.Number, invoice.Amount, invoice.Currency, ID)
	invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, msg))
}
