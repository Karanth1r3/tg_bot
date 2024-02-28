package commands

import (
	"github.com/Karanth1r3/tg_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, inputMessage *tgbotapi.Message){}

type Commander struct {
	bot         *tgbotapi.BotAPI
	productSrvc *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, psrvc product.Service) *Commander {
	return &Commander{
		bot:         bot,
		productSrvc: &psrvc,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {

	if update.Message == nil { // If we got a message
		return
	}

	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.DefaultBehaviour(update.Message)
	}

	switch update.Message.Command() {
	case "help":
		c.HelpCommand(update.Message)
	case "list":
		c.ListCommand(update.Message)
	default:
		c.DefaultBehaviour(update.Message)
	}
}
