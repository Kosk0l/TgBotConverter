package handlers

import (
	"strings"

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

func (h *Handler) HandleUpdate(update telegram.Update) {
	if update.Message == nil {
		return
	}

	// TODO: реализовать проверку наличия пользователя
	// h.u.GetByIdService()

	if update.Message.IsCommand() {
		h.HandleCommand(update)
		return
	}	

	if update.Message.Document != nil {
		h.HandleDocument(update)
		return
	}

	h.HandleText(update)
}

//====================================================================================================

func (h *Handler) HandleCommand(update telegram.Update) {
	chatID := update.Message.Chat.ID

	switch update.Message.Command() {
	case "start":
		h.bot.Send(telegram.NewMessage(chatID,"Привет! Отправь документ для конвертации."))
	default:
		h.bot.Send(telegram.NewMessage(chatID,"Неизвестная команда"))
	}
}

func (h *Handler) HandleDocument(update telegram.Update) {
	doc := update.Message.Document
	chatID := update.Message.Chat.ID
	filename := strings.ToLower(doc.FileName)

	switch {
	case strings.HasSuffix(filename, ".pdf"):
		h.handlePDF(update)

	case strings.HasSuffix(filename, ".docx"):
		h.handleDOCX(update)

	case strings.HasSuffix(filename, ".xlsx"):
		h.handleXLSX(update)

	default:
		h.bot.Send(telegram.NewMessage(chatID, "Этот тип файла пока не поддерживается"))
	}
}
