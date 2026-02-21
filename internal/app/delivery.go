package app

import (
	"log/slog"

	converterworker "github.com/Kosk0l/TgBotConverter/internal/ConverterWorker"
	"github.com/Kosk0l/TgBotConverter/internal/handlers"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func initDelivery(bot *telegram.BotAPI, s *Services, l *slog.Logger) (*handlers.Handler, *converterworker.Worker) {
    handler := handlers.NewServer(bot, s.User, s.Job, s.Dialog, l)
    worker := converterworker.NewWorker(s.Job, s.Converter, l)

    return handler, worker
}