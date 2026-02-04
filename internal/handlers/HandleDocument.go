package handlers

import (
	"context"
	"log/slog"

	"github.com/Kosk0l/TgBotConverter/internal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Хендлер документов
func (h *Handler) HandleDocument(ctx context.Context, update telegram.Update) {
	h.log.Info("user send document",
		slog.Int64("chat_id", update.Message.Chat.ID),
	)

	// Получить fileUrl
	file := update.Message.Document
	fileUrl, err := h.bot.GetFileDirectURL(file.FileID)
	if err != nil {
		h.log.Error("error - get file in handler",
			slog.Int64("chat_id", update.Message.Chat.ID),
			slog.Any("error", err),
		)
		return
	}

	// создать состояние
	state := domains.State{
		ChatId: update.Message.Chat.ID,
		Step: domains.WaitingTargetType,
		FileURL: fileUrl,
		FileName: file.FileName,
		Size: int64(file.FileSize),
		ContentType: file.MimeType,
	}

	// Бизнес-логика - добавить состояние
	if err := h.ds.SetState(ctx, state); err != nil {
		h.log.Error("error - set state in handler",
			slog.Int64("chat_id", update.Message.Chat.ID),
			slog.Any("error", err),
		)
		return
	}

	// Подключение кнопок
	msg := telegram.NewMessage(update.Message.Chat.ID, "В какой тип необходимо преобразовать?")
	msg.ReplyMarkup = targetTypeKeyboard()
	h.bot.Send(msg)

	h.log.Info("Success processing - set state",
		slog.Int64("chat_id", update.Message.Chat.ID),
		slog.String("file_url", state.FileURL),
	)
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