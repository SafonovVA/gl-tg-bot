package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	TelegramBotToken string
	EnableTLS        string
	WebhookUrl       string
}

func NewConfig() *config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	return &config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		EnableTLS:        os.Getenv("ENABLE_TLS"),
		WebhookUrl:       os.Getenv("WEBHOOK_URL"),
	}
}
