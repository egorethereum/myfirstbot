package main

import (
	"./Database"
	"MyFirstBot/Configs"
	_ "fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("7"),
		tgbotapi.NewKeyboardButton("8"),
		tgbotapi.NewKeyboardButton("9"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("+"),
		tgbotapi.NewKeyboardButton("-"),
		tgbotapi.NewKeyboardButton("x"),
		tgbotapi.NewKeyboardButton("="),
	),
)

func main() {

	bot, err := tgbotapi.NewBotAPI(Configs.Token)
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

		var response string
		if update.Message.Text == "/start"{
           Database.AddUser ( update.Message.From.FirstName, update.Message.From.LastName, strconv.Itoa(update.Message.From.ID), update.Message.From.UserName)
           response = " User successfully appended to database"
		}

		if update.Message.Text == "/firstname" {
			response = update.Message.From.FirstName

		} else if update.Message.Text == "/lastname" {
			response = update.Message.From.LastName
		} else if update.Message.Text == "/nick" {
			response = (update.Message.From.FirstName) + " " + (update.Message.From.LastName)
		} /*else if update.Message.Text == "/showme"{
			Database.ShowUser(update.Message.From.ID)
		}*/


		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyMarkup = numericKeyboard

		bot.Send(msg)
	}
}
