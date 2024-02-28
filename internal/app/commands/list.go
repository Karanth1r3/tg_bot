package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) ListCommand(inputMessage *tgbotapi.Message) {

	outputMsgText := "All products: \n\n"

	products := c.productSrvc.Toist()

	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	c.bot.Send(msg)
}
