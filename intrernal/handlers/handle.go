package handlers

import (
	"context"

	Dialogservice "github.com/Kosk0l/TgBotConverter/intrernal/Services/DialogService"
	jobservice "github.com/Kosk0l/TgBotConverter/intrernal/Services/jobService"
	userService "github.com/Kosk0l/TgBotConverter/intrernal/Services/userService"
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

//====================================================================================================

// TODO: Дальше можно разрезать по зонам ответственности: ht *HandlerText
type Handler struct {
	bot *telegram.BotAPI
	us 	*userService.UserService
	js 	*jobservice.JobService
	ds 	*Dialogservice.DialogService
}

// Конструктор
func NewServer(bot *telegram.BotAPI, us *userService.UserService, js *jobservice.JobService) (*Handler) {
	return &Handler{
		bot: bot,
		us: us,
		js: js,
	}
}

//====================================================================================================

// Распределяет по типам сообщения
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

