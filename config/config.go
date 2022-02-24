package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	TelegramBotToken string
}

func NewConfig() *config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	return &config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}
