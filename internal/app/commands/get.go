package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) GetCommand(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong arg", err, args)
		return
	}

	outputMsg := ""

	product, err := c.productSrvc.Get(arg)
	if err != nil {
		log.Printf("failed to get product with idx %d : %v", arg, err)
		outputMsg = "failed to get product, " + err.Error()
	} else {
		outputMsg = "Requested product: " + product.Title
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
	c.bot.Send(msg)
}
