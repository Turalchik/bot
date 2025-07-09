package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/task-1/internal/model/payment"
)

func (invoiceService *DummyInvoiceService) Describe(invoiceID uint64) (*payment.Invoice, error) {
	invoicesIndex, ok := invoiceService.ID2InvoicesIndex[invoiceID]
	if !ok {
		return nil, fmt.Errorf("non-existent ID: %v", invoiceID)
	}
	return &invoiceService.Invoices[invoicesIndex], nil
}
