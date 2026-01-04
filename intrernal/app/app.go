package app // Связка компонентов

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Kosk0l/TgBotConverter/config"
	"github.com/Kosk0l/TgBotConverter/intrernal/handlers"
	"github.com/Kosk0l/TgBotConverter/intrernal/storage"
	"github.com/Kosk0l/TgBotConverter/intrernal/userService"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct { 
	bot *telegram.BotAPI
	handler *handlers.Handler
}

// Конструктор
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	// объект бот
	bot, err := telegram.NewBotAPI(cfg.App.TOKEN)
	if err != nil {
		return nil, fmt.Errorf("error in up telegram token(newapp constructor): %v", err)
	}

	// объект постгреса 
	dsn := config.LoadDsn(cfg)
	pool, err := storage.NewPostgres(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("error in up storage: %v", err)
	}
	
	// Объект сервиса
	userService := userservice.NewService(pool)

	// объект хендлера
	handler := handlers.NewServer(bot, userService)

	bot.Debug = true
	log.Printf("\nAuthorized on account %s", bot.Self.UserName)
	
	return &App {
		bot: bot,
		handler: handler,
	}, nil
}

func (a *App) Run(ctx context.Context) () {
	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30

	// канал чтения из апи тг
	updates := a.bot.GetUpdatesChan(u)

	// проходка по каналу
	for update := range updates {
		go func(update telegram.Update) {
			ctxUpdate, cancel := context.WithTimeout(ctx, 15*time.Second)
			defer cancel() // выход из горутины
			a.handler.HandleUpdate(ctxUpdate, update)
		}(update)
	}
}

func (a *App) Stop() () {
	a.bot.StopReceivingUpdates()
	log.Println("telegram bot stopped")
}