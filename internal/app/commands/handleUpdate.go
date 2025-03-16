package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander *Commander) HandleUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		commander.Help(update.Message)
	case "list":
		commander.List(update.Message)
	case "get":
		commander.Get(update.Message)
	default:
		commander.Default(update.Message)
	}
}
