package sender

import "github.com/Turalchik/bot/internal/model/payment"

type EventSender interface {
	Send(invoice *payment.InvoiceEvent) error
}
