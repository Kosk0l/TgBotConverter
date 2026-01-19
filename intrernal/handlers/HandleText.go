package handlers

import(
	"context"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработчик текстов
func (h *Handler) HandleText(ctx context.Context, update telegram.Update) {
	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID,"Я принимаю только команды и документы"))
}
