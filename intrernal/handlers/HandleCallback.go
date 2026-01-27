package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) HandleCallBack(ctx context.Context, update telegram.Update) () {
	cb := update.CallbackQuery
	chatId := cb.Message.Chat.ID

	// Получим состояние
	state, err := h.ds.GetState(ctx, chatId) 
	if err != nil {
		log.Printf("handler - failed getstate service: %v", err)
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
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка обработки - файл не поддерживается"))
		h.bot.Send(telegram.NewMessage(chatId,"Неизвестный формат файла"))
		return
	}

	// Создание JobService
	jobId, err := h.js.CreateJob(ctx, job, obj)
	if err != nil {
		log.Printf("handler - failed create job: %v", err)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка - невозможно добавить дальше"))
		h.bot.Send(telegram.NewMessage(chatId,"Недостаточно данных"))
		return
	}

	// Отправка ответов
	h.bot.Request(telegram.NewCallback(cb.ID, fmt.Sprintf("%s успешно", jobId)))
	h.bot.Send(telegram.NewMessage(chatId,"Добавили в очередь выполнения"))
}