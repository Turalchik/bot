package invoice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (invoiceCommander *DummyInvoiceCommander) handleMessage(inputMsg *tgbotapi.Message) {
	switch command := inputMsg.Command(); command {
	case "get__payment__invoice":
		invoiceCommander.Get(inputMsg)
	case "list__payment__invoice":
		invoiceCommander.List(inputMsg)
	case "delete__payment__invoice":
		invoiceCommander.Delete(inputMsg)
	case "edit__payment__invoice":
		invoiceCommander.Edit(inputMsg)
	case "new__payment__invoice":
		invoiceCommander.New(inputMsg)
	default:
		invoiceCommander.Help(inputMsg)
	}
}
