package commands

import (
	"fmt"
	"log"
	"strings"

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

	defer func() {
		if panicVal := recover(); panicVal != nil {
			log.Printf("recovered from panic: %v", panicVal)
		}
	}()

	// Ignore any non-Message update
	if upd.Message == nil {
		return
	}

	if upd.CallbackQuery != nil {
		args := strings.Split(upd.CallbackQuery.Data, "_")
		msg := tgbotapi.NewMessage(
			upd.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %s\n", args[0])+
				fmt.Sprintf("Offset: %s\n", args[1]),
		)
		c.bot.Send(msg)
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
