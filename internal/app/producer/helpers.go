package producer

import "github.com/Turalchik/bot/internal/model/payment"

func chanInvoiceEvents2SliceInvoiceEventsANDIDs(ch <-chan payment.InvoiceEvent) ([]payment.InvoiceEvent, []uint64) {
	eventIDs := make([]uint64, 0, len(ch))
	events := make([]payment.InvoiceEvent, 0, len(ch))

	for {
		select {
		case event := <-ch:
			eventIDs = append(eventIDs, event.ID)
			events = append(events, event)
		default:
			return events, eventIDs
		}
	}
}
