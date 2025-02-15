package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cm *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "Here is all products:\n\n"
	products := cm.productService.List()

	for _, product := range products {
		outputMessage += product.Title + "\n"
	}

	serializedData, _ := json.Marshal(CommandData{Offset: 10})

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	button := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)
	msg.ReplyMarkup = button

	cm.bot.Send(msg)
}
