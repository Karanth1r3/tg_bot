package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Karanth1r3/tg_bot/internal/app/commands"
	"github.com/Karanth1r3/tg_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	fmt.Println(token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	//Initializing services
	productSrvc := product.NewService()

	commander := commands.NewCommander(bot, *productSrvc)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.HelpCommand(update.Message)
		case "list":
			commander.ListCommand(update.Message)
		default:
			commander.DefaultBehaviour(update.Message)
		}

	}
}
