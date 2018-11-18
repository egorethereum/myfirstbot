package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var response string

		switch update.Message.Text {
		case "/start":
			response = "Hello!"
		default:
			response = "Not hello bitch!"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID,response)

		bot.Send(msg)
	}
}