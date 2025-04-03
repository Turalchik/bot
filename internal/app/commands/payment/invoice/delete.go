package invoice

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"unicode"
)

func (invoiceCommander *DummyInvoiceCommander) Delete(inputMsg *tgbotapi.Message) {
	arg := strings.TrimFunc(inputMsg.CommandArguments(), unicode.IsSpace)
	if arg == "" {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! Формат: /delete__payment__invoice <ID>"))
		return
	}

	ID, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка! ID должен быть натуральным числом"))
		return
	}

	isDeleted, err := invoiceCommander.invoiceService.Remove(ID)
	outText := "Счёт удалить не удалось\n"
	if isDeleted {
		outText = "Счёт успешно удалён\n"
	}
	if err != nil {
		outText += err.Error() + "\n"
	}

	invoiceCommander.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, outText))
}
