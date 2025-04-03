package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/internal/model/payment"
	"strconv"
)

func getInvoice(args []string) (*payment.Invoice, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("ошибка! Формат ввода счёта: <номер> <сумма> <валюта>")
	}

	amount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("ошибка! Сумма должна быть числом")
	}

	return &payment.Invoice{
		Number:   args[0],
		Amount:   amount,
		Currency: args[2],
	}, nil
}
