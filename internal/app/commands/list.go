package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander *Commander) List(inputMessage *tgbotapi.Message) {
	outMsgTxt := "Вот вам лист товаров\n\n"

	for _, prod := range commander.productService.List() {
		outMsgTxt += prod.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMsgTxt)
	commander.bot.Send(msg)
}
