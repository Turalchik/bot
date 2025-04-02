package payment

type Invoice struct {
	Number   string  `json:"number"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
