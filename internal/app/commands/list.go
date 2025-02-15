package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (cm *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "Here is all products:\n\n"
	products := cm.productService.List()

	for _, product := range products {
		outputMessage += product.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	cm.bot.Send(msg)
}
