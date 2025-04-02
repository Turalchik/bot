package invoice

import "github.com/Turalchik/bot/internal/model/payment"

type InvoiceService interface {
	Describe(invoiceID uint64) (*payment.Invoice, error)
	List(cursor uint64, limit uint64) ([]payment.Invoice, error)
	Create(payment.Invoice) (uint64, error)
	Update(invoiceID uint64, invoice payment.Invoice) error
	Remove(invoiceID uint64) (bool, error)
}

type DummyInvoiceService struct {
	ID2Invoice map[uint64]*payment.Invoice
}

func NewDummyInvoiceService() *DummyInvoiceService {
	return &DummyInvoiceService{
		ID2Invoice: make(map[uint64]*payment.Invoice),
	}
}
