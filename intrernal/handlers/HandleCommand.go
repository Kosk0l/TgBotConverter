package handlers

import (
	"context"
	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Хендлер команд
func (h *Handler) HandleCommand(ctx context.Context, update telegram.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start":
		// Проверка наличия пользователя
		_, err := h.us.GetByIdService(ctx, update.Message.From.ID)
		if err != nil {
			// Создание модели
			var user domains.User
			user.ID = update.Message.From.ID
			user.UserName = update.Message.From.UserName
			user.FirstName = update.Message.From.FirstName
			user.LastName = update.Message.From.LastName

			// Создание пользователя
			h.us.CreateUserService(ctx, &user)
		}
		h.bot.Send(telegram.NewMessage(chatID,"Привет! Отправь документ для конвертации."))
	default:
		h.bot.Send(telegram.NewMessage(chatID,"Неизвестная команда"))
	}
}