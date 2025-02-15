package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cm *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args: ", args)
		return
	}

	product, err := cm.productService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with ids %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)

	cm.bot.Send(msg)
}
