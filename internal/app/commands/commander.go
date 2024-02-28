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
