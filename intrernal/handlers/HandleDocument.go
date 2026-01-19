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
	chatID := update.Message.Chat.ID
	fileUrl, err := h.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		log.Printf("handler - failed get file url: %v", err)
		return
	}

	State := domains.State{
		FileURL: fileUrl,
		FileName: file.FileName,
		ChatId: chatID,
		Size: int64(file.FileSize),
		ContentType: file.MimeType,
	}

	

	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID,"В какой тип необходимо преобразовать?"))
}