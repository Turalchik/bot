package payment

type Invoice struct {
	ID       uint64  `json:"id"`
	Number   string  `json:"number"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
