package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (commander *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	num, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("Invalid args: %s", args)
		return
	}

	outMsgTxt := "Число чётное"
	if num%2 != 0 {
		outMsgTxt = "Число нечётное"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMsgTxt)
	commander.bot.Send(msg)
}
