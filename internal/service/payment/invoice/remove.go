package invoice

import (
	"fmt"
)

func (invoiceService *DummyInvoiceService) Remove(invoiceID uint64) (bool, error) {
	invoicesIndex, ok := invoiceService.ID2InvoicesIndex[invoiceID]
	if !ok {
		return false, fmt.Errorf("non-existent ID: %v", invoiceID)
	}

	delete(invoiceService.ID2InvoicesIndex, invoiceID)
	invoiceService.Size--
	invoiceService.Invoices[invoicesIndex] = invoiceService.Invoices[invoiceService.Size]

	return true, nil
}
