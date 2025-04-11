package payment

type EventStatus uint8
type EventType uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type InvoiceEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Invoice
}
