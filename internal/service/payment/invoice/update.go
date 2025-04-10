package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/internal/model/payment"
)

func (invoiceService *DummyInvoiceService) Update(invoiceID uint64, invoice payment.Invoice) error {
	invoicesIndex, ok := invoiceService.ID2InvoicesIndex[invoiceID]
	if !ok {
		return fmt.Errorf("non-existent ID")
	}
	invoiceService.Invoices[invoicesIndex] = invoice
	invoiceService.Invoices[invoicesIndex].ID = invoiceID
	return nil
}
