package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gl-tg-bot/config"
	"log"
	"net/http"
)

func main() {
	configs := config.NewConfig()
	bot, err := tgbotapi.NewBotAPI(configs.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+bot.Token, tgbotapi.FilePath("cert.pem"))

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go func() {
		err := http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		//_, err := bot.Send(tgbotapi.NewMessage(-703603408, update.Message.Text))
		//if err != nil {
		//	log.Fatal("SUKA")
		//}
		log.Printf("%+v\n", update)
	}
}
