package handlers

import (
	"context"
	"strings"

	"github.com/Kosk0l/TgBotConverter/intrernal/models"
	userService "github.com/Kosk0l/TgBotConverter/intrernal/userService"
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

func (h *Handler) HandleCommand(ctx context.Context, update telegram.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start":
		h.bot.Send(telegram.NewMessage(chatID,"Привет! Отправь документ для конвертации."))
		_, err := h.u.GetByIdService(ctx, update.Message.From.ID)
		if err != nil {
			var user models.User
			user.ID = update.Message.From.ID
			user.UserName = update.Message.From.UserName
			user.FirstName = update.Message.From.FirstName
			user.LastName = update.Message.From.LastName
			h.u.CreateUserService(ctx, &user)
		}

	default:
		h.bot.Send(telegram.NewMessage(chatID,"Неизвестная команда"))
	}
}

func (h *Handler) HandleDocument(ctx context.Context, update telegram.Update) {
	doc := update.Message.Document
	chatID := update.Message.Chat.ID
	filename := strings.ToLower(doc.FileName)

	switch {
	case strings.HasSuffix(filename, ".pdf"):
		h.handlePDF(ctx, update)

	case strings.HasSuffix(filename, ".docx"):
		h.handleDOCX(ctx, update)

	case strings.HasSuffix(filename, ".xlsx"):
		h.handleXLSX(ctx, update)

	default:
		h.bot.Send(telegram.NewMessage(chatID, "Этот тип файла пока не поддерживается"))
	}
}
