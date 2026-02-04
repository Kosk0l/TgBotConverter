package app // Связка компонентов

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Kosk0l/TgBotConverter/config"
	converterworker "github.com/Kosk0l/TgBotConverter/internal/ConverterWorker"
	converterservice "github.com/Kosk0l/TgBotConverter/internal/Services/ConverterService"
	Dialogservice "github.com/Kosk0l/TgBotConverter/internal/Services/DialogService"
	jobservice "github.com/Kosk0l/TgBotConverter/internal/Services/jobService"
	"github.com/Kosk0l/TgBotConverter/internal/Services/userService"
	"github.com/Kosk0l/TgBotConverter/internal/handlers"
	"github.com/Kosk0l/TgBotConverter/internal/lib/logger"
	"github.com/Kosk0l/TgBotConverter/internal/storage/cache"
	"github.com/Kosk0l/TgBotConverter/internal/storage/minio"
	"github.com/Kosk0l/TgBotConverter/internal/storage/postgres"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct { 
	bot 	*telegram.BotAPI
	handler *handlers.Handler
	worker 	*converterworker.Worker
}

// Конструктор
func NewApp(ctx context.Context, cfg config.Config) (*App, error) {
	// объект бота телеграм
	bot, err := telegram.NewBotAPI(cfg.App.TOKEN)
	if err != nil {
		return nil, fmt.Errorf("error in up telegram token(newapp constructor): %w", err)
	}

	// Объект логгера
	logger := logger.NewLogger(cfg)

	// объект постгреса 
	dsn := config.LoadDsn(cfg)
	pool, err := postgres.NewPostgres(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("error in up storage: %w", err)
	}

	// объект редиса
	cache, err := cache.NewRedis(ctx, cfg) 
	if err != nil {
		return nil, fmt.Errorf("error in up redis: %w", err)
	}

	// объект минио
	minio, err := minio.NewMinio(ctx, cfg, "files") 
	if err != nil {
		return nil ,fmt.Errorf("error in up minio: %w", err)
	}
	
	// Объекты сервисов
	userService := userservice.NewUserService(pool)
	jobService := jobservice.NewJobService(cache, minio)
	dialogService := Dialogservice.NewDialogService(cache)
	converterservice := converterservice.NewConverterService()

	// объект хендлера
	handler := handlers.NewServer(bot, userService, jobService, dialogService, logger) 

	// объект воркера
	worker := converterworker.NewWorker(jobService, converterservice)

	bot.Debug = true
	log.Printf("\nAuthorized on account %s", bot.Self.UserName)
	
	return &App {
		bot: bot,
		handler: handler,
		worker: worker,
	}, nil
}

func (a *App) Run(ctx context.Context) () {
	// запуск воркера
	go a.worker.Run(ctx)

	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30
	u.AllowedUpdates = []string{"message", "callback_query"}

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