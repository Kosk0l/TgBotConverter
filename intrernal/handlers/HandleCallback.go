package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) HandleCallBack(ctx context.Context, update telegram.Update) () {
	cb := update.CallbackQuery
	chatId := cb.Message.Chat.ID

	// Логирование нажатия
	log.Printf("Callback received: %s", cb.Data)

	// Получим состояние
	state, err := h.ds.GetState(ctx, chatId) 
	if err != nil {
		log.Printf("handler - failed getstate service: %v", err)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка"))
		return
	}

	job := domains.Job{
		ChatID: chatId,
	}

	data, err := http.Get(state.FileURL)
	if err != nil {
		log.Printf("handler - failed http get file: %v", err)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка"))
		return
	}
	defer data.Body.Close()

	obj := domains.Object{
		Reader: data.Body,
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
		h.bot.Request(telegram.NewCallback(cb.ID, "Неизвестный формат на данный момент"))
		return
	}

	jobId, err := h.js.CreateJob(ctx, job, obj)
	if err != nil {
		log.Printf("handler - failed create job: %v", err)
		h.bot.Request(telegram.NewCallback(cb.ID, "Ошибка"))
		return
	}

	h.bot.Request(telegram.NewCallback(cb.ID, fmt.Sprintf("%s успешно", jobId)))
	h.bot.Send(telegram.NewMessage(chatId,"Добавили в очередь выполнения"))
}