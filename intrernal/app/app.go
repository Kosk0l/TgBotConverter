package app // Связка компонентов

import (
	"log"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/handlers"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct { 
}

func NewApp(cfg *config.Config) () {
	bot, err := telegram.NewBotAPI(cfg.App.TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30

	// канал чтения из апи тг
	updates := bot.GetUpdatesChan(u)

	// проходка по каналу
	for update := range updates {
		handlers.HandleUpdate(bot, update)
	}
}