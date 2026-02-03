package handlers

import (
	"context"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
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
			//TODO: ввести error.Is и на него накинуть условие и выполнение

			// Создание domain
			user := domains.User{
				ID: update.Message.From.ID,
				UserName: update.Message.From.UserName,
				FirstName: update.Message.From.FirstName,
				LastName: update.Message.From.LastName,
			}

			// Создание пользователя
			err := h.us.CreateUserService(ctx, user) 
			if err != nil {
				h.bot.Send(telegram.NewMessage(chatID,"На данный момент сервис недоступен"))
				return
			}
		}
		h.bot.Send(telegram.NewMessage(chatID,"Привет! Отправь документ для конвертации."))
	default:
		h.bot.Send(telegram.NewMessage(chatID,"Неизвестная команда"))
	}
}