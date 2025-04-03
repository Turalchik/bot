package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) HandleUpdate(update *tgbotapi.Update) {
	switch command := update.Message.Command(); command {
	case "get__payment__invoice":
		invoiceCommander.Get(update.Message)
	case "list__payment__invoice":
		invoiceCommander.List(update.Message)
	case "delete__payment__invoice":
		invoiceCommander.Delete(update.Message)
	case "edit__payment__invoice":
		invoiceCommander.Edit(update.Message)
	case "new__payment__invoice":
		invoiceCommander.New(update.Message)
	default:
		invoiceCommander.Help(update.Message)
	}
}
