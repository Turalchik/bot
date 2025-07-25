package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/task-1/internal/model/payment"
)

func (invoiceService *DummyInvoiceService) Create(invoice payment.Invoice) (uint64, error) {
	if invoice.Amount < 0 {
		return 0, fmt.Errorf("invalid fields of invoice")
	}

	invoiceService.ID2InvoicesIndex[invoiceService.FreeID] = invoiceService.Size
	invoiceService.Invoices = append(invoiceService.Invoices, invoice)
	invoiceService.Invoices[invoiceService.Size].ID = invoiceService.FreeID
	invoiceService.Size++
	invoiceService.FreeID++

	return invoiceService.FreeID - 1, nil
}
