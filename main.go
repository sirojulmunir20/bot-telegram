package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := "5813596879:AAE74N1qnv2gl0TXj8wHgxGjDikiIYlxrk8"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if strings.Contains(update.Message.Text, "ready") {
				message, err := os.ReadFile("template/ready.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
	
				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "harga"){
				message,err := os.ReadFile("template/harga.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else {
				welcomeMessage := fmt.Sprintf("Hallo %s, Selamat datang di toko kita kak 🥳🥳 \n [📝] Silahkan order...", update.Message.From.UserName) 
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeMessage))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
