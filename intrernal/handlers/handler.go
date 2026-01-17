package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"

	userService "github.com/Kosk0l/TgBotConverter/intrernal/Services/userService"
	"github.com/Kosk0l/TgBotConverter/intrernal/models"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// bot - http cliet
// update - http request // Содержит всю информацию

//
type Handler struct {
	bot *telegram.BotAPI
	u 	*userService.UserService
}

func NewServer(bot *telegram.BotAPI, u *userService.UserService) (*Handler) {
	return &Handler{
		bot: bot,
		u: u,
	}
}

//====================================================================================================
// ОБЩИЙ ОБРАБОТЧИК

func (h *Handler) HandleUpdate(ctx context.Context, update telegram.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		h.HandleCommand(ctx, update)
		return
	}	

	if update.Message.Document != nil {
		h.HandleDocument(ctx, update)
		return
	}

	h.HandleText(ctx, update)
}

//====================================================================================================

// Хендлер команд
func (h *Handler) HandleCommand(ctx context.Context, update telegram.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start":
		// Проверка наличия пользователя
		_, err := h.u.GetByIdService(ctx, update.Message.From.ID)
		if err != nil {
			// Создание модели
			var user models.User
			user.ID = update.Message.From.ID
			user.UserName = update.Message.From.UserName
			user.FirstName = update.Message.From.FirstName
			user.LastName = update.Message.From.LastName

			// Создание пользователя
			h.u.CreateUserService(ctx, &user)
		}
		h.bot.Send(telegram.NewMessage(chatID,"Привет! Отправь документ для конвертации."))
	default:
		h.bot.Send(telegram.NewMessage(chatID,"Неизвестная команда"))
	}
}

// Хендлер документов
func (h *Handler) HandleDocument(ctx context.Context, update telegram.Update) {
	// Получить file
	file := update.Message.Document
	chatID := update.Message.Chat.ID
	fileUrl, _ := h.bot.GetFileDirectURL(file.FileID)

	// стрим байтов
	resp, err := http.Get(fileUrl) 
	if err != nil {
		log.Printf("handler - failed get file: %s", err)
		return 
	}
	defer resp.Body.Close() // Закрыть поток данных

	switch {
	case strings.HasSuffix(file.FileName, ".pdf"):
		h.handlePDF(ctx, update)

	case strings.HasSuffix(file.FileName, ".docx"):
		h.handleDOCX(ctx, update)

	case strings.HasSuffix(file.FileName, ".xlsx"):
		h.handleXLSX(ctx, update)

	default:
		h.bot.Send(telegram.NewMessage(chatID, "Этот тип файла пока не поддерживается"))
	}
}
