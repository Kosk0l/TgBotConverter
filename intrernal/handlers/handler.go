package handlers

import (
	"context"
	"log"

	jobservice "github.com/Kosk0l/TgBotConverter/intrernal/Services/jobService"
	userService "github.com/Kosk0l/TgBotConverter/intrernal/Services/userService"
	"github.com/Kosk0l/TgBotConverter/intrernal/domains"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// bot - http cliet
// update - http request // Содержит всю информацию

// Вид json запроса: (все есть в update)

	/*
	"message":{
		"message_id":19,
		"from":{
			"id":7792217214,
			"is_bot":false,
			"first_name":"Николай",
			"username":"kosk0l",
			"language_code":"ru"
		},
		"chat":{
			"id":7792217214,
			"first_name":"Николай",
			"username":"kosk0l",
			"type":"private"
		},
		"date":1766682293,
		"text":"В"
	}
	*/

// TODO: Дальше можно разрезать по зонам ответственности: ht *HandlerText
type Handler struct {
	bot *telegram.BotAPI
	us 	*userService.UserService
	js 	*jobservice.JobService
}

func NewServer(bot *telegram.BotAPI, us *userService.UserService, js *jobservice.JobService) (*Handler) {
	return &Handler{
		bot: bot,
		us: us,
		js: js,
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

	InquiryJob := domains.State{
		FileURL: fileUrl,
		FileName: file.FileName,
		ChatId: chatID,
		Size: int64(file.FileSize),
		ContentType: file.MimeType,
	}

	

	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID,"В какой тип необходимо преобразовать?"))
}

//====================================================================================================

// Обработчик текстов
func (h *Handler) HandleText(ctx context.Context, update telegram.Update) {
	h.bot.Send(telegram.NewMessage(update.Message.Chat.ID,"Я принимаю только команды и документы"))
}

