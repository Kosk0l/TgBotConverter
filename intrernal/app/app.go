package app // Связка компонентов

import (
	"context"
	"log"
	"time"

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

	// dsn := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	cfg.Db.User,
	// 	cfg.Db.Pass,
	// 	cfg.Db.Host,
	// 	cfg.Db.Port,
	// 	cfg.Db.Name,
	// )

	//TODO: реализовать конструктор pgxpool
	
	// Объект сервиса
	userservice := userservice.NewService(nil)

	// объект хендлера
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
		go func(update telegram.Update) {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			a.handler.HandleUpdate(ctx, update)
		}(update)
	}
}

func (a *App) Stop() () {

}