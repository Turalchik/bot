package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) Help(inputMsg *tgbotapi.Message) {
	command2Description := map[string]string{
		"/get__payment__invoice <ID>":                           "Возвращает информацию о счете по его идентификатору",
		"/list__payment__invoice <limit>":                       "Получает список счетов с поддержкой постраничной навигации",
		"/delete__payment__invoice <ID>":                        "Удаляет счет по идентификатору и возвращает статус операции",
		"/edit__payment__invoice <ID> <номер> <сумма> <валюта>": "Обновляет данные существующего счета",
		"/new__payment__invoice <номер> <сумма> <валюта>":       "Создает новый счет и возвращает его идентификатор",
	}

	outText := "Список всех комманд, которые вы можете использовать:\n\n"
	for command, description := range command2Description {
		outText += command + " - " + description + "\n"
	}

	outMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, outText)
	invoiceCommander.bot.Send(outMessage)
}
