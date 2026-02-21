package app // Связка компонентов

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Kosk0l/TgBotConverter/config"
	converterworker "github.com/Kosk0l/TgBotConverter/internal/ConverterWorker"
	"github.com/Kosk0l/TgBotConverter/internal/handlers"
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

	Infrastructure, err := initInfrastructure(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("new app error - :%w", err)
	}
	
	Services := initServices(Infrastructure)

	handler, worker := initDelivery(bot, Services, Infrastructure.Logger)

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
	workersCount := 5
	for i := 0; i < workersCount; i++ {
		go a.worker.Run(ctx)
	}

	// Настраиваем получение апдейтов
	u := telegram.NewUpdate(0)
	u.Timeout = 30
	u.AllowedUpdates = []string{"message", "callback_query"}

	// канал чтения из апи тг
	updates := a.bot.GetUpdatesChan(u)

	// Семафор хендлера
	handlerPoll := make(chan struct{}, 100)

	// проходка по каналу
	for update := range updates {
		handlerPoll <- struct{}{}
		go func(update telegram.Update) {
			defer func() {
				<- handlerPoll
			}()
			ctxUpdate, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel() // выход из горутины
			a.handler.HandleUpdate(ctxUpdate, update)
		}(update)
	}
}

func (a *App) Stop() () {
	a.bot.StopReceivingUpdates()
	log.Println("telegram bot stopped")
}