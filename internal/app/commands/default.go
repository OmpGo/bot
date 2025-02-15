package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cm *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text)

	cm.bot.Send(msg)
}

func (cm *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.Message == nil { // If we got a message
		return
	}

	switch update.Message.Command() {
	case "help":
		cm.Help(update.Message)
	case "list":
		cm.List(update.Message)
	case "get":
		cm.Get(update.Message)
	default:
		cm.Default(update.Message)
	}
}
