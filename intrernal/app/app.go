package app // Связка компонентов

import (
	"log"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/handlers"
	userservice "github.com/Kosk0l/TgBotConverter/intrernal/userService"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct { 
	bot *telegram.BotAPI
	handler *handlers.Handler
}

// Конструктор
func NewApp(cfg *config.Config) (*App) {
	bot, err := telegram.NewBotAPI(cfg.App.TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: реализовать конструктор pgxpool
	
	userservice := userservice.NewService(nil)

	handler := handlers.NewServer(bot, userservice)

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	
	return &App {
		bot: bot,
		handler: handler,
	}
}

func (a *App) Run() () {
	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30

	// канал чтения из апи тг
	updates := a.bot.GetUpdatesChan(u)

	// проходка по каналу
	for update := range updates {
		a.handler.HandleUpdate(update)
	}
}