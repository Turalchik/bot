package sender

import "github.com/Turalchik/bot/internal/model/payment"

type EventSender interface {
	Send(subdomain *payment.InvoiceEvent) error
}
