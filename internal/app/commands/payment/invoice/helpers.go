package invoice

import (
	"fmt"
	"github.com/Turalchik/bot/internal/model/payment"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func firstRowForKeyboardMarkup(cursor uint64, limit uint64, numberInvoices uint64) []tgbotapi.InlineKeyboardButton {
	if cursor == 0 && cursor+limit >= numberInvoices {
		return tgbotapi.NewInlineKeyboardRow()
	}
	if cursor > 0 && cursor+limit >= numberInvoices {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<- предыдущие счета", "previous_invoices"),
		)
	}
	if cursor == 0 && cursor+limit < numberInvoices {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("следующие счета ->", "next_invoices"),
		)
	}

	return tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<- предыдущие счета", "previous_invoices"),
		tgbotapi.NewInlineKeyboardButtonData("следующие счета ->", "next_invoices"),
	)
}

func arrOfInvoices2Txt(arrOfInvoices []payment.Invoice) string {
	var outTxt string
	for _, invoice := range arrOfInvoices {
		outTxt += fmt.Sprintf("%s %.2f %s\n", invoice.Number, invoice.Amount, invoice.Currency)
	}
	return outTxt
}
