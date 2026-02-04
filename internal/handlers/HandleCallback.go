package handlers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Kosk0l/TgBotConverter/internal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) HandleCallBack(ctx context.Context, update telegram.Update) () {
	cb := update.CallbackQuery
	chatId := cb.Message.Chat.ID

	h.log.Info("user send the callback",
		slog.Int64("chat_id", chatId),
	)

	// Получим состояние
	state, err := h.ds.GetState(ctx, chatId) 
	if err != nil {
		h.log.Error("error - get state in handler",
			slog.Int64("chat_id", chatId),
			slog.Any("error", err),
		)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка - вы еще не добавили файл"))
		return
	}

	// Создание domains
	job := domains.Job{
		ChatID: chatId,
	}

	obj := domains.Object{
		FlieURL: state.FileURL,
		Size: state.Size,
		ContentType: state.ContentType,
	}

	// выбор конвертации
	switch cb.Data {
	case "to:pdf":
		job.FileTypeTo = domains.Pdf
	case "to:docx":
		job.FileTypeTo = domains.Docx
	case "to:jpeg":
		job.FileTypeTo = domains.Jpeg
	case "to:xlsx":
		job.FileTypeTo = domains.Xlsx
	default:
		h.log.Error("error - the file is not implemented",
			slog.Int64("chat_id", chatId),
			slog.Any("error", err),
		)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка обработки - невозможная реализация"))
		h.bot.Send(telegram.NewMessage(chatId,"Неизвестный формат файла"))
		return
	}

	// Создание JobService
	jobId, err := h.js.CreateJob(ctx, job, obj)
	if err != nil {
		h.log.Error("error - create job in handler",
			slog.Int64("chat_id", chatId),
			slog.Any("error", err),
		)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка - невозможно добавить дальше"))
		h.bot.Send(telegram.NewMessage(chatId,"Недостаточно данных"))
		return
	}

	// Отправка ответов
	h.bot.Request(telegram.NewCallback(cb.ID, fmt.Sprintf("%s успешно", jobId)))
	h.bot.Send(telegram.NewMessage(chatId,"Добавили в очередь выполнения"))
	h.log.Info("success create the job",
		slog.Int64("chat_id", chatId),
		slog.String("file_url", state.FileURL),
	)
}