package commands

import (
	"github.com/Karanth1r3/tg_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func (c *Commander) HandleUpdate(upd tgbotapi.Update) {
	// Ignore any non-Message update
	if upd.Message == nil {
		return
	}

	switch upd.Message.Command() {
	case "help":
		c.HelpCommand(upd.Message)
	case "list":
		c.ListCommand(upd.Message)
	case "get":
		c.GetCommand(upd.Message)
	default:
		c.DefaultBehaviour(upd.Message)
	}
}
