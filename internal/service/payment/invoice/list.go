package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/internal/model/payment"
)

func (invoiceService *DummyInvoiceService) List(cursor uint64, limit uint64) ([]payment.Invoice, error) {
	if cursor >= invoiceService.Size {
		return nil, fmt.Errorf("there are too few invoices for this cursor\n")
	}
	return invoiceService.Invoices[cursor:min(cursor+limit, invoiceService.Size)], nil
}
