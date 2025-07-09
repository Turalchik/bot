package invoice

import (
	"github.com/Turalchik/bot/task-1/internal/model/payment"
)

type InvoiceService interface {
	Describe(invoiceID uint64) (*payment.Invoice, error)
	List(cursor uint64, limit uint64) ([]payment.Invoice, error)
	Create(payment.Invoice) (uint64, error)
	Update(invoiceID uint64, invoice payment.Invoice) error
	Remove(invoiceID uint64) (bool, error)
	GetSize() uint64
}

type DummyInvoiceService struct {
	Invoices         []payment.Invoice
	FreeID           uint64
	ID2InvoicesIndex map[uint64]uint64
	Size             uint64
}

func NewDummyInvoiceService() *DummyInvoiceService {
	return &DummyInvoiceService{
		FreeID:           1,
		ID2InvoicesIndex: make(map[uint64]uint64),
	}
}
