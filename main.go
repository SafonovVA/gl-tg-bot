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

	var wh tgbotapi.WebhookConfig
	if configs.EnableTLS == "enable" {
		wh, _ = tgbotapi.NewWebhookWithCert("https://www.example.com:8443/"+bot.Token, tgbotapi.FilePath("cert.pem"))
	} else {
		wh, _ = tgbotapi.NewWebhook("https://www.example.com:8443/" + bot.Token)
	}

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
		var err error
		if configs.EnableTLS == "enable" {
			err = http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
		} else {
			err = http.ListenAndServe("0.0.0.0:8443", nil)
		}
		if err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
