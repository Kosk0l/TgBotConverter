package main

import (
	log "log"
	os "os"

	"github.com/Kosk0l/TgBotConverter/intrernal/handlers"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}

	// Создаём бота
	bot, err := telegram.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true // можно отключить потом
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handlers.HandleUpdate(bot, update)
	}
}