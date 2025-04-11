package repo

import "github.com/Turalchik/bot/internal/model/payment"

type EventRepo interface {
	Lock(n uint64) ([]payment.InvoiceEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []payment.InvoiceEvent) error
	Remove(eventIDs []uint64) error
}
