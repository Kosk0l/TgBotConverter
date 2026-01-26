package handlers

import (
	"context"
	"log"
	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


// Хендлер документов
func (h *Handler) HandleDocument(ctx context.Context, update telegram.Update) {

	// Получить file
	file := update.Message.Document
	fileUrl, err := h.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		log.Printf("handler - failed get file url: %v", err)
		return
	}

	// создать состояние
	state := domains.State{
		ChatId: update.Message.Chat.ID,
		UserId: update.Message.From.ID,
		Step: domains.WaitingTargetType,
		FileURL: fileUrl,
		FileName: file.FileName,
		Size: int64(file.FileSize),
		ContentType: file.MimeType,
	}

	// Бизнес-логика - добавить состояние
	if err := h.ds.SetState(ctx, state); err != nil {
		log.Printf("handler - failed setstate service: %v", err)
		return
	}

	// Подключение кнопок
	msg := telegram.NewMessage(update.Message.Chat.ID,"В какой тип необходимо преобразовать?")
	msg.ReplyMarkup = targetTypeKeyboard()
	h.bot.Send(msg)
}

// Функция - добавление кнопок 
func targetTypeKeyboard() telegram.InlineKeyboardMarkup {
	return telegram.NewInlineKeyboardMarkup(
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData("📄 PDF", "to:pdf"),
			telegram.NewInlineKeyboardButtonData("📝 DOCX", "to:docx"),
		),
		telegram.NewInlineKeyboardRow(
			telegram.NewInlineKeyboardButtonData("📊 XLSX", "to:xlsx"),
		),
	)
}